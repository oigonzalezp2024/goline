package main

import (
	"fmt"
	"log"
	"net/http"
)

func logConsole() {
	fmt.Println("Petición recibida en /api/hello v1")
}

func sendMessage(w http.ResponseWriter) {
	fmt.Fprintln(w, "Hello, World!!!!")
}

func registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		logConsole()
		sendMessage(w)
	})
}

func main() {
	mux := http.NewServeMux()
	registerRoutes(mux)

	fmt.Println("Servidor escuchando en http://localhost:8080")

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
