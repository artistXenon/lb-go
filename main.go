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
var compiled_regexp, _ = regexp.Compile(`^(([-a-zA-Z0-9.]+)\.|)uxagee\.in$`)

func onHttps(res http.ResponseWriter, req *http.Request) {
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

func runHttps(cfg *configs.Config) {
	http.Handle("/", http.HandlerFunc(onHttps))
	err := http.ListenAndServeTLS(fmt.Sprintf(":%d", cfg.SecurePort), cfg.Certificate, cfg.PrivateKey, nil)

	if err != nil {
		log.Fatal("Serve: ", err)
	}
}

func onHttp(res http.ResponseWriter, _ *http.Request) {
	res.WriteHeader(http.StatusBadRequest)
	res.Write([]byte("Insecure access, Redirect manually"))
}

func runHttp(port uint16) {
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	if err != nil {
		log.Fatal("Serve: ", err)
	}
}

func main() {
	config := configs.Load()

	go runHttps(config)
	go runHttp(config.Port)

	<-make(chan int)
}
