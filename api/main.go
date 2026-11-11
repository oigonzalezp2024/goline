package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Registramos la ruta
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		// Este mensaje solo saldrá en la consola cuando recibas una petición HTTP
		fmt.Println("Petición recibida en /api/hello")
		fmt.Fprintln(w, "Hello, World!")
	})

	// Mensaje para confirmar en consola que el servidor arrancó
	fmt.Println("Servidor escuchando en http://localhost:8080")

	// ListenAndServe bloquea la ejecución aquí. Si falla, Log.Fatal mostrará el error.
	log.Fatal(http.ListenAndServe(":8080", nil))
}
