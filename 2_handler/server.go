package handler

import (
	"net/http"
)

type Server struct {
	http.Handler
}

func NewServer() *Server {
	srv := &Server{}

	router := http.NewServeMux()
	router.HandleFunc("/participants", Register)
	router.HandleFunc("/participants/", Show)

	srv.Handler = router
	return srv
}
