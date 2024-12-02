package main

import (
	"log"
	"net/http"
)

// Manejador de la ruta /users
func userHandler(w http.ResponseWriter, r *http.Request) {
	// Aquí escribirías el código para manejar las solicitudes
	log.Println("Manejando solicitud para /users")
	w.Write([]byte("Usuarios"))
}
