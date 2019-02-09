package headr

import (
	"net/http"
	"strings"
)

// Headr contains the definition of the HTTP handler decorator
type Headr struct {
	handler http.Handler
	headers [][]string
}

func (h *Headr) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, header := range h.headers {
		w.Header().Set(header[0], header[1])
	}
	h.handler.ServeHTTP(w, req)
}

// New creates a new HTTP handler decorator
func New(h http.Handler) *Headr {
	return &Headr{handler: h, headers: [][]string{}}
}

// Set adds header entries with the given key to the HTTP handler decorator
func (h *Headr) Set(key string, values []string) {
	value := strings.Join(values, ",")
	header := []string{key, value}
	h.headers = append(h.headers, header)
}
