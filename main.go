package main

import (
	"fmt"
	"log"
	"net/http"
)

// Manejador para la ruta raíz
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido a la API de Go!")
}

func main() {
	// Registra el manejador para la ruta raíz
	http.HandleFunc("/", homeHandler)

	// Inicia el servidor HTTP
	log.Fatal(http.ListenAndServe(":8080", nil))
}
