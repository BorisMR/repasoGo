package main

import (
	"fmt"
	"gomysql/db"
	"gomysql/models"
)

func main() {
	db.Conectar()
	db.Ping()

	// cargar tabla si esta no existe
	db.CargarTabla(models.UserSchema, "usuarios")
	// limpia la tabla si esta tiene datos
	//db.TruncateTable("usuarios")

	/*
		listamos usuarios
		Agregamos un usuario
		lo editamos
		listamos usuarios
		lo eliminamos
		listamos nuevamente
	*/
	usrsLista := models.ListarUsuarios()
	fmt.Println(usrsLista)

	testUSR := models.NuevoUsuario("Boris", "lapass", "anum.bmr.01@gmail.com")
	fmt.Println(testUSR)
	testUSR.UserName = "Sirob"
	testUSR.Guardar()
	fmt.Println(testUSR)

	usrsLista = models.ListarUsuarios()
	fmt.Println(usrsLista)

	testUSR.Eliminar()
	usrsLista = models.ListarUsuarios()
	fmt.Println(usrsLista)

	db.Cerrar()
}
