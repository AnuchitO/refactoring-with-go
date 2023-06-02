package main

import (
	"context"
	"fmt"
)

func main() {
	var word string
	fmt.Println("What do you want to say?")
	s := &CliSay{}
	fmt.Scan(&word)
	s.Saying(context.Background(), word)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "form.html")
	// })

	// http.HandleFunc("/say", func(w http.ResponseWriter, r *http.Request) {
	// 	r.ParseForm()
	// 	word := r.FormValue("word")
	// 	w.Write([]byte(Saying(word)))
	// })
	// http.ListenAndServe("localhost:8080", nil)
}

type CliSay struct{}

func (c *CliSay) Saying(ctx context.Context, word string) {
	fmt.Printf("Say, %s!\n", word)
}
