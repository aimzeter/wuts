package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Server struct {
	http.Handler
}

func NewServer() *Server {
	srv := &Server{}

	router := httprouter.New()
	router.GET("/participants/:nik", Show)
	router.POST("/participants", Register)

	srv.Handler = router
	return srv
}
