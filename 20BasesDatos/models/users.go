package models

type User struct {
	Id       int
	UserName string
	Password string
	Email    string
}

const UserSchema string = `CREATE TABLE usuarios (
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	nombre VARCHAR(6) NOT NULL,
	contrasenna VARCHAR(100) NOT NULL,
	email VARCHAR(50),
	created_data TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`
