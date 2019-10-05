package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	//"io/ioutil"
	"net"
)

func main() {

	ua := "{{ .VARIABLE }}" //"https://example.com\uFF03bing.com"

	u, err := url.Parse(ua)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("URL: %v\n", u.String())
	fmt.Printf("Scheme: %v\n", u.Scheme)
	fmt.Printf("Opaque: %v\n", u.Opaque)
	fmt.Printf("User: %v\n", u.User)
	fmt.Printf("Host: %v\n", u.Host)
	fmt.Printf("Hostname(): %v\n", u.Hostname())
	fmt.Printf("Path: %v\n", u.Path)
	fmt.Printf("RawPath: %v\n", u.RawPath)
	fmt.Printf("RawQuery: %v\n", u.RawQuery)
	fmt.Printf("Fragment: %v\n", u.Fragment)

	conn, err := net.Dial("tcp", ":80")

	fmt.Printf("%v", conn)
	resp, err := http.Get("http://" + u.Host)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	/*
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatal(err)
		}
	*/
}
