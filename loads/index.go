package loads

import (
	"errors"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

var m = map[string]http.Handler{}

func GetServe(sub_domain string) (handler http.Handler) {
	h := getDirectory(sub_domain)
	if h != nil {
		return h
	}
	h = getPipe(sub_domain)
	if h != nil {
		return h
	}
	return nil
}

func getDirectory(sub_domain string) (handler http.Handler) {
	path := "./serve/" + sub_domain
	if _, err := os.Stat(path); err == nil {
		val, exists := m[sub_domain]
		if !exists {
			val = http.FileServer(http.Dir(path))
			m[sub_domain] = val
		}
		return val
	} else if errors.Is(err, os.ErrNotExist) {
		return nil
	} else {
		return nil
	}
}

func getPipe(sub_domain string) (handler http.Handler) {
	// TODO: config
	// TODO: remove unused handler
	if sub_domain == "workspace" {
		val, exists := m[sub_domain]
		if !exists {
			remote, err := url.Parse("http://192.168.0.1")
			if err != nil {
				panic(err)
			}
			proxy := httputil.NewSingleHostReverseProxy(remote)
			val = &ProxyHandler{proxy}
			m[sub_domain] = val
		}
		// TODO: create a fallback
		// dest server may not be open. in this case, return a generated page about it
		return val
	}
	return nil
}

// TODO: clear map.
// every n minutes, check connection usage and clear reference to specific http handler so that GC can do a job.
// make map will save not just handler but a struct with such info
