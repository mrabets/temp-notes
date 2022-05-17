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
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method is not allowed", 405)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(`{ title: "Lorem ipsum" }`))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/note", showNote)
	mux.HandleFunc("/note/create", createNote)

	log.Println("Server is running on http://127.0.0.1:9090")
	log.Fatal(http.ListenAndServe(":9090", mux))
}
