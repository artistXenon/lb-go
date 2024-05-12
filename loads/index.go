package loads

import (
	"errors"
	"net/http"
	"os"
)

var m map[string]http.Handler

func GetServe(sub_domain string) (handler http.Handler) {
	path := "./serve/" + sub_domain
	if _, err := os.Stat(path); err == nil {
		val, exists := m[sub_domain]
		if !exists {
			val = http.FileServer(http.Dir("path"))
			m[sub_domain] = val
		}
		return val
	} else if errors.Is(err, os.ErrNotExist) {
		return nil
	} else {
		return nil
	}
}
