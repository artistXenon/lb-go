package main

import (
	"fmt"

	configs "lb-go/configs"
	loads "lb-go/loads"
	"log"
	"net/http"
	"regexp"
)

// parse based on config
var compiled_regexp, _ = regexp.Compile(`^(([-a-zA-Z0-9.]+)\.|)jaewon\.pro$`)

func onHttp(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	hostname := req.Host

	fmt.Printf("%s %s\n", hostname, path)

	sub_match := compiled_regexp.FindStringSubmatch(hostname)

	if len(sub_match) == 3 {
		sub_domain := sub_match[2]
		if sub_domain == "" {
			// root handler
			res.Write([]byte("Coming Soon :)"))
			return
		}

		handler := loads.GetServe(sub_domain)
		if handler != nil {
			handler.ServeHTTP(res, req)
			return
		}
	}

	res.WriteHeader(http.StatusInternalServerError)
	res.Write([]byte("Internal Error"))
}

func main() {
	config := configs.Load()

	http.Handle("/", http.HandlerFunc(onHttp))
//	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
	err := http.ListenAndServeTLS(fmt.Sprintf(":%d", config.SecurePort), "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal("Serve: ", err)
	}
}
