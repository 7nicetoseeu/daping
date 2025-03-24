package httphandler

import (
	"fmt"
	"net/http"
)

func writeError(w http.ResponseWriter, error string, code int) {
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	fmt.Fprintln(w, error)
}

type notFound struct{}

func (self *notFound) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", req.Header.Get("Origin"))
	w.Header().Add("Content-Type", req.Header.Get("Content-Type"))

	writeError(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

type methodNotAllow struct{}

func (self *methodNotAllow) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", req.Header.Get("Origin"))
	w.Header().Add("Content-Type", req.Header.Get("Content-Type"))

	writeError(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}
