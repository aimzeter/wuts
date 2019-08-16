package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		return
	}

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

func Show(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(405)
		return
	}

	nik := r.URL.Path[len("/participants/"):]
	fmt.Fprintf(w, "Showing participant with NIK %s", nik)
}
