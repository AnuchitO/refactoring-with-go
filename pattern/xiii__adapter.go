package main

import "net/http"

type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

var serve = &ServeMux{}

type ServeMux struct {
}

func (mux *ServeMux) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	mux.Handle(pattern, HandlerFunc(handler))
}

func (mux *ServeMux) Handle(pattern string, handler Handler) {

}

type HandlerFunc func(http.ResponseWriter, *http.Request)

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}

func Handle(pattern string, handler http.Handler) {
	serve.Handle(pattern, handler)
}

func HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	serve.HandleFunc(pattern, handler)
}

func main() {
	Handle("/path", HandlerFunc(func(http.ResponseWriter, *http.Request) {
		// ...
	}))

	HandleFunc("/path", func(writer http.ResponseWriter, request *http.Request) {
		// ...
	})
}
