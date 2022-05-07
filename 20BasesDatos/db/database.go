package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//username:password@tcp(localhost:3306)/database

const url_db = "root:gol3m....gol3m@tcp(localhost:3306)/golang_db"

var db *sql.DB

func Conectar() {
	conn, err := sql.Open("mysql", url_db)

	if err != nil {
		panic("")
	}

	fmt.Println("¡Conexión Abierta!")

	db = conn
}

func Cerrar() {
	fmt.Println("¡Conexión Cerrada!")
	db.Close()
}

func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("¡Conexión Sigue Abierta!")
	}
}

func CargarSchemma(schema string, tablename string) {
	if !ExisteTabla(tablename) {
		fmt.Println("¡Cargando Esquema!")
		result, _ := db.Exec(schema)
		fmt.Println(result)
		fmt.Println("¡Esquema Cargado!")
	} else {
		fmt.Println("¡La tabla ya se encuentra cargada!")
	}
}

func ExisteTabla(tablename string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tablename)
	rows, err := db.Query(sql)

	if err != nil {
		fmt.Println("ERROR! => ", err)
	}

	return rows.Next()
}
