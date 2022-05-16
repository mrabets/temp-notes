package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello World"))
}

func showNote(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Current note"))
}

func createNote(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create new note"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/note", showNote)
	mux.HandleFunc("/note/create", createNote)

	log.Println("Server is running on http://127.0.0.1:9090")
	log.Fatal(http.ListenAndServe(":9090", mux))
}
