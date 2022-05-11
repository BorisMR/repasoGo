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

func CargarTabla(schema string, tableName string) {
	if !ExisteTabla(tableName) {
		result, err := db.Exec(schema)

		if err != nil {
			fmt.Println("¡Error en Esquema!")
			fmt.Println(err)
		} else {
			fmt.Println("¡Esquema Cargado!")
			fmt.Println(result)
		}

	} else {
		fmt.Println("¡La Tabla Ya Se Encuentra Cargada!")
	}
}

func ExisteTabla(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := Query(sql)

	if err != nil {
		fmt.Println("ERROR! => ", err)
	}

	return rows.Next()
}

// limpiar datos de tabla, reinicia Id
func TruncateTable(tableName string) {
	sql := fmt.Sprintf("TRUNCATE %s", tableName)
	Exec(sql)

	fmt.Println("->Limpiando Tabla", tableName)
}

// polimorfizando Exec y Query
func Exec(qry string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(qry, args...)

	if err != nil {
		fmt.Println(err)
	}

	return result, err
}

func Query(qry string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(qry, args...)

	if err != nil {
		fmt.Println(err)
	}

	return rows, err
}
