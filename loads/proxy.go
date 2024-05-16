package loads

import (
	"net/http"
	"net/http/httputil"
)

type ProxyHandler struct {
	p *httputil.ReverseProxy
}

func (ph *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ph.p.ServeHTTP(w, r)
}
