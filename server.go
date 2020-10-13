package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "you have invoked /hello endpoint")
}

func main() {
	/*
		http.HandleFunc("/hello", helloHandler)
		// log.Print(http.ListenAndServe(":8080", nil))
		log.Print(http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil))
	*/

	caCert, err := ioutil.ReadFile("cert.pem")
	if err != nil {
		log.Fatal(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := tls.Config{
		ClientCAs:  caCertPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}

	tlsConfig.BuildNameToCertificate()

	server := &http.Server{
		Addr:      ":8080",
		TLSConfig: &tlsConfig,
	}
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(server.ListenAndServeTLS("cert.pem", "key.pem"))
}
