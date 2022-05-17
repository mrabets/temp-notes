package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/note", showNote)
	mux.HandleFunc("/note/create", createNote)

	log.Println("Server is running on http://127.0.0.1:9090")
	log.Fatal(http.ListenAndServe(":9090", mux))
}
