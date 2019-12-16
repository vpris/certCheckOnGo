package certCheckonGo

import (
	"crypto/tls"
	"flag"
	"log"
	"net/url"
	"time"
)

func main() {

	serverFlag := flag.String("servername", "example.com", "a string")

	flag.Parse()

	servername := *serverFlag

	u, err := url.Parse(servername)

	if err != nil {
		log.Fatal(err)
	}
	// Connect to site
	cfg := tls.Config{}
	conn, err := tls.Dial("tcp", u.Host+":443", &cfg)
	if err != nil {
		log.Fatalln("TLS connection failed: " + err.Error())
	}

	certChain := conn.ConnectionState().PeerCertificates

	for _, crt := range certChain {

		if len(crt.DNSNames) > 0 {
			notAfte := crt.NotAfter
			notBefo := crt.NotBefore
			d := notAfte.Sub(notBefo)
			daysBalance := notAfte.Sub(time.Now())
			d = time.Duration(d.Hours() / 24)
			daysBalance = time.Duration(daysBalance.Hours() / 24)
			percent := (daysBalance * 100) / d
			print(percent)
		}
	}
}
