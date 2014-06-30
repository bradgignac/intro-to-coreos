package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Printf("Received request for %s", req.URL.Path)
		io.WriteString(w, "Hello Go!\n")
	})

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
