package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	var resp *http.Response
	var client *http.Client

	var host string
	var err error
	var showHeaders bool

	flag.StringVar(&host, "h", "", "TLS servername to use in preference to URL host")
	flag.BoolVar(&showHeaders, "i", false, "show request/response headers")
	flag.Parse()

	url := flag.Arg(0)
	if url == "" {
		fmt.Println("You must supply a URL")
		return
	}

	req, _ := http.NewRequest("GET", url, nil)
	transport := &http.Transport{
		DisableKeepAlives:  true,
		DisableCompression: true,
	}
	if host != "" {
		// Set hostname for TLS connection. This allows us to connect using
		// another hostname or IP for the actual TCP connection. Handy for GeoDNS scenarios.
		transport.TLSClientConfig = &tls.Config{
			ServerName: host,
		}
		req.Host = host
	}
	client = &http.Client{
		Transport: transport,
	}
	resp, err = client.Do(req)
	if err != nil {
		log.Printf("http(s) error, %s", err)
	} else {
		if showHeaders {
			// Go doesn't populate request header for client
			//fmt.Printf("\nRequest:\n\n%s", req.Header)
			header := ""
			for key, _ := range resp.Header {
				header += key + ": " + resp.Header.Get(key) + "\n"
			}
			fmt.Printf("Response Header:\n------------------\n%s\n", header)
		}
		var body []byte
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("http(s) error, %s", err)
		} else {
			fmt.Printf("%s", body)
		}
		resp.Body.Close()
	}
}
