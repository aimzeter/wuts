package dukcapil_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

type DukcapilHandler struct {
	Endpoint string
	Body     string
	Code     int
}

func SetupDukcapilServer(h *DukcapilHandler) *httptest.Server {
	router := http.NewServeMux()
	router.HandleFunc(h.Endpoint, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(h.Code)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, h.Body)
	}))

	return httptest.NewServer(router)
}

func (h *DukcapilHandler) Stub(body string, code int) {
	h.Body = body
	h.Code = code
}
