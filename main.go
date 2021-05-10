package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type HomeHellow struct {
	Message string
}

func main() {

	r := chi.NewRouter()

	r.Get("/", HomePage)

	http.ListenAndServe(":3000", r)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	res := HomeHellow{
		Message: "WHALEOCME HERE",
	}
	json.NewEncoder(w).Encode(res)
}
