package models

import (
	"gomysql/db"
)

type User struct {
	Id       int64
	UserName string
	Password string
	Email    string
}

type Users []User

const UserSchema string = `CREATE TABLE usuarios (
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	nombre VARCHAR(6) NOT NULL,
	contrasenna VARCHAR(100) NOT NULL,
	email VARCHAR(50),
	created_data TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`

// crear ususario usando struct
func NuevoUsuario(nombre, contrasenna, email string) *User {
	usuario := &User{
		UserName: nombre,
		Password: contrasenna,
		Email:    email,
	}

	return usuario
}

// define funcion para insertar usuario
func (usr *User) Insertar() {
	sql := "INSERT usuarios SET nombre=?, contrasenna=?, email=?"
	result, _ := db.Exec(sql, usr.UserName, usr.Password, usr.Email)
	usr.Id, _ = result.LastInsertId()
}

// crea un ususario y lo inserta en la base de datos
func CrearUsuario(nombre, contrasenna, email string) *User {
	usr := NuevoUsuario(nombre, contrasenna, email)
	usr.Insertar()

	return usr
}

func ListarUsuarios() Users {
	sql := "SELECT id, nombre, contrasenna, email FROM usuarios"
	listaUsuarios := Users{}
	rows, _ := db.Query(sql)

	for rows.Next() {
		usuario := User{}
		rows.Scan(&usuario.Id, &usuario.UserName, &usuario.Password, &usuario.Email)
		listaUsuarios = append(listaUsuarios, usuario)
	}

	return listaUsuarios
}

// obtener por id
func ObtenerUsuarioById(id int) *User {
	usr := NuevoUsuario("", "", "")
	sql := "SELECT id, nombre, contrasenna, email FROM usuarios WHERE id=?"
	rows, _ := db.Query(sql, id)

	for rows.Next() {
		rows.Scan(&usr.Id, &usr.UserName, &usr.Password, &usr.Email)
	}

	return usr
}

//Actualizar Registro
func (usr *User) Actualizar() {
	sql := "UPDATE usuarios SET nombre=?, contrasenna=?, email=? WHERE id=?"
	db.Exec(sql, usr.UserName, usr.Password, usr.Email, usr.Id)
}

// Guardar un registro
func (usr *User) Guardar() {
	if usr.Id == 0 {
		usr.Insertar()
	} else {
		usr.Actualizar()
	}
}

// eliminar
func (usr *User) Eliminar() {
	sql := "DELETE FROM usuarios WHERE id=?"
	db.Exec(sql, usr.Id)
}
