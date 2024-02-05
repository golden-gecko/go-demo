package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	/*
	fmt.Println(primitive.ObjectID{})
	*/

	/*
	for i := 0; i < 100; i++ {
		t0 := time.Now().UTC()
		t1 := t0.Format("2006-01-02 15:04:05.000000")

		fmt.Println(t1)
	}
	*/

	url := "https://api.com.gecko:4000/v1/healthcheck1"
	caCert, err := os.ReadFile("../certs/ca-cert.pem")

	if err != nil {
        log.Fatal(err)
    }

    caCertPool := x509.NewCertPool()
    caCertPool.AppendCertsFromPEM(caCert)

    client := &http.Client{
        Transport: &http.Transport{
            TLSClientConfig: &tls.Config{
                RootCAs:      caCertPool,
            },
        },
    }

	log.Info(url)

	response, err := client.Get(url)

    if err != nil {
        log.Fatal(err)
    }

	body, err := io.ReadAll(response.Body)

    if err != nil {
		log.Info(fmt.Sprintf("%d %s", response.StatusCode, string(body)))
    } else {
		log.Info(response.StatusCode)
	}
}
