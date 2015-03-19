package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	var resp *http.Response
	var client *http.Client

	var host string
	var err error
	var showHeaders bool

	flag.StringVar(&host, "h", "", "website hostname to use in preference to URL hostname")
	flag.BoolVar(&showHeaders, "i", false, "show response headers")
	flag.Parse()

	url := flag.Arg(0)
	if url == "" {
		fmt.Fprintln(os.Stderr, "You must supply a URL")
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
		fmt.Fprintf(os.Stderr, "http(s) error, %s", err)
	} else {
		if showHeaders {
			// Go doesn't populate request header for client
			//fmt.Printf("\nRequest:\n\n%s", req.Header)
			header := ""
			for key, _ := range resp.Header {
				header += fmt.Sprintf("%20s: %s\n", key, resp.Header.Get(key))
			}
			fmt.Fprintf(os.Stderr, "%20s\n", "Response Headers")
			fmt.Fprintf(os.Stderr, "%20s\n", "-----------------")
			fmt.Fprintf(os.Stderr, "%s\n", header)
		}
		var body []byte
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "http(s) error, %s", err)
		} else {
			fmt.Printf("%s", body)
		}
		resp.Body.Close()
	}
}
