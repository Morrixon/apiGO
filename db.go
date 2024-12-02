package main

import (
"database/sql"
"log"

_"github.com/mattn/go-sqlite3"
)

//Conectar a la base de datos de SQLITE
func getDBconn() *sql.DB {
//abre la base de datos "users.db"
db , err :=sql.Open("sqlite3", "./users.db")
if err != nil {
	//Si hay error al abrir la base de datos , lo muestra y termina el programa
	log.Fatal("Error al abrir la base de datos: ", err)
}
return db
}

//creaTable crea la tabla de usuario si no existe

func creaTable(db *sql.DB) {
	//Sql para crear la tabla "users"
	createSQL := `
	CREATE TABLE IF NOT EXIST users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	email TEXT NOT NULL
	);`
}

// Ejecuta la consulta SQL
_, err := db.Exec(createSQL)
if err != nil {
	// Si hay error muestra lo siguiente
	log.Fatal("Error al crear la tabla de usuarios: ", err)
}