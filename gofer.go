package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	// make sure we have at least one command line argument
	if len(os.Args) <= 1 {
		fmt.Println("Please provide a URL and port to test (ie. brave.com:443")
		return
	}

	// loop over all command line arguments
	for _, arg := range os.Args[1:] {
		url := arg
		server := strings.Split(arg, ":")

		conn, err := net.Dial("tcp", url)
		if err != nil {
			log.Fatal(err)
		}

		client := tls.Client(conn, &tls.Config{
			ServerName: (server[0]),
		})
		defer client.Close()

		if err := client.Handshake(); err != nil {
			log.Fatal(err)
		}

		cert := client.ConnectionState().PeerCertificates[0]
		fmt.Println("-- ", server[0], " --")
		fmt.Println("Certificate:")
		fmt.Println("\tDNS Name:", cert.DNSNames)
		fmt.Println("\tSerial Number:", cert.SerialNumber)
		fmt.Println("\tIssuer:\t", cert.Issuer)
		fmt.Println("\tValidity:")
		fmt.Println("\t\tNot Before:\t", cert.NotBefore)
		fmt.Println("\t\tNot After:\t", cert.NotAfter)
		fmt.Println("\tSubject:", cert.Subject)
		fmt.Println("\tAuthority Information Access:")
		fmt.Println("\t\t", "OCSP - URI:\t", cert.OCSPServer)
		fmt.Println("\t\t", "CA Issuers - URI:\t", cert.IssuingCertificateURL)
		fmt.Println("\tSignatureAlgorithm:", cert.SignatureAlgorithm)
		fmt.Println("\tPublicKeyAlgorithm:", cert.PublicKeyAlgorithm)
		fmt.Println("--\n")

	}
}
