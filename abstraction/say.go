package main

import (
	"net/http"
)

func main() {
	// var word string
	// fmt.Println("What do you want to say?")
	// fmt.Scan(&word)
	// fmt.Printf("Say, %s!\n", word)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "form.html")
	})

	http.HandleFunc("/say", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		word := r.FormValue("word")
		w.Write([]byte("Say, " + word + "!\n"))
	})
	http.ListenAndServe("localhost:8080", nil)
}
