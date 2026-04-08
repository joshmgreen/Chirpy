package main

import (
	"log"
	"net/http"
)

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome home!"))
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir(".")))

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
