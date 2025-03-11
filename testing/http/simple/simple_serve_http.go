package simple

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", SimpleHandler)

	log.Fatal(http.ListenAndServe(":1234", nil))
}

func SimpleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"name": "anuchit"}`))
}
