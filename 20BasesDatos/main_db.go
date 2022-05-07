package main

import (
	"gomysql/db"
	"gomysql/models"
)

func main() {
	db.Conectar()
	db.CargarSchemma(models.UserSchema, "usuarios")
	db.Cerrar()
}
