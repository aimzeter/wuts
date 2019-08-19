package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var regBody struct {
		NIK   string `json:"nik"`
		Agree bool   `json:"agree"`
	}

	err := json.NewDecoder(r.Body).Decode(&regBody)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(201)
	fmt.Fprintf(w, "Participant with NIK %s successfully registered", regBody.NIK)
}

func Show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	nik := ps.ByName("nik")
	fmt.Fprintf(w, "Showing participant with NIK %s", nik)
}
