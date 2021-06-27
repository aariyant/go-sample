package main

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

//API function
func API(method, endpoint string) []byte {

	//BACKEND Configuration
	URL := os.Getenv("URL")
	expiration := time.Duration(5 * time.Second)

	// TODO: This is insecure; use only in dev environments.
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := http.Client{
		Timeout:   expiration,
		Transport: tr,
	}

	req, err := http.NewRequest(method, URL+endpoint, nil)
	if err != nil {
		log.Println("error on creating new request", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("error on creating request", err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("error on reading request", err)
	}

	log.Println(string(b))

	return b

}
