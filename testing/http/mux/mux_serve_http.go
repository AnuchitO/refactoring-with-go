package mux

import (
	"log"
	"net/http"
)

func MuxHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"name": "anuchit"}`))
}

func NewServeMux() *http.ServeMux {
	m := http.NewServeMux()

	m.HandleFunc("/", MuxHandler)

	return m
}

func main() {
	m := NewServeMux()

	s := &http.Server{
		Addr:    ":8000",
		Handler: m,
	}

	log.Fatal(s.ListenAndServe())
}
