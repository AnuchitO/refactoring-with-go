package main

import (
	"context"
	"fmt"
	"net/http"
)

func main() {
	// var word string
	// fmt.Println("What do you want to say?")
	// s := &CliSay{}
	// fmt.Scan(&word)
	// s.Saying(context.Background(), word)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "form.html")
	})

	h := Handler{
		s: &WebSay{},
	}

	http.HandleFunc("/say", h.SayingHandler)
	http.ListenAndServe("localhost:8080", nil)
}

type Say interface {
	Saying(ctx context.Context, word string)
}

type Handler struct {
	s Say
}

type writerType string

const writerKey writerType = "writer"

func (h *Handler) SayingHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	word := r.FormValue("word")
	ctx := context.WithValue(r.Context(), writerKey, w)
	h.s.Saying(ctx, word)
}

type WebSay struct {
}

func (w *WebSay) Saying(ctx context.Context, word string) {
	wt := ctx.Value(writerKey).(http.ResponseWriter)
	wt.Write([]byte("Say, " + word + "!"))
}

type CliSay struct{}

func (c *CliSay) Saying(ctx context.Context, word string) {
	fmt.Printf("Say, %s!\n", word)
}
