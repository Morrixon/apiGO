package main

import (
	"database/sql"
	"log" // Asegúrate de que este import esté aquí si deseas usarlo
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestCreateTable(t *testing.T) {
	// Crea una base de datos en memoria para las pruebas (sin archivo en disco)
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal("Error al abrir la base de datos en memoria:", err) // Usamos log.Fatal
	}
	defer db.Close()

	// Llamamos a la función que crea la tabla
	CreateTable(db)

	// Verificamos que la tabla "users" existe
	row := db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='users';")
	var tableName string
	err = row.Scan(&tableName)
	if err != nil {
		log.Fatal("Error al verificar si la tabla 'users' fue creada:", err) // Usamos log.Fatal
	}
	if tableName != "users" {
		t.Errorf("Se esperaba que la tabla 'users' estuviera creada, pero no se encontró")
	}
}
