package di

import (
	"fmt"
	"log"
	"net/http"
)

type Email interface {
	Send(msg string)
}

func DIMuxHandler(email Email) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := `{"name": "anuchit"}`
		email.Send(msg)
		w.Write([]byte(msg))
	}

}

func NewServeMuxDI(email Email) *http.ServeMux {
	m := http.NewServeMux()

	m.HandleFunc("/", DIMuxHandler(email))

	return m
}

type email struct {
}

func (email) Send(msg string) {
	fmt.Println("real sent email: ", msg)
}

func main() {
	m := NewServeMuxDI(&email{})

	s := &http.Server{
		Addr:    ":8000",
		Handler: m,
	}

	log.Fatal(s.ListenAndServe())
}
