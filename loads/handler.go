package loads

import "net/http"

type IHandler interface {
	GetHandler() http.Handler
}
