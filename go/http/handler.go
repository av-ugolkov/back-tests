package http

import (
	"log"
	"net/http"
)

func helloGoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Go!"))
}

func Start() {
	http.HandleFunc("/", helloGoHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
