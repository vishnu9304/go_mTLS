package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	/*
			resp, err := http.Get("http://localhost:8080/hello")
			if err != nil {
				log.Fatal(err)
			}

		defer resp.Body.Close()

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(data))
	*/

	cert, err := tls.LoadX509KeyPair("/Users/v01/MyCodeRepos/github/goMTLS/server/cert.pem", "/Users/v01/MyCodeRepos/github/goMTLS/server/key.pem")
	if err != nil {
		log.Fatal(err)
	}

	caCert, err := ioutil.ReadFile("/Users/v01/MyCodeRepos/github/goMTLS/server/cert.pem")
	if err != nil {
		log.Fatal(err)
	}

	caCertPool := x509.NewCertPool()

	caCertPool.AppendCertsFromPEM(caCert)

	tlsClientConfig := tls.Config{
		RootCAs:      caCertPool,
		Certificates: []tls.Certificate{cert},
	}

	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tlsClientConfig,
		},
	}

	resp, err := client.Get("https://localhost:8080/hello")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
