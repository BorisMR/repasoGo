package main

import "fmt"

// realcion 1 a 1
type Usuario struct {
	nombre string
	email  string
	activo bool
}
type Estudiante struct {
	usuario Usuario
	codigo  string
}

// relacion 1 a *
type Curso struct {
	titulo string
	videos []Video
}
type Video struct {
	titulo string
	curso  Curso
}

func main() {
	// probando relacion 1 a 1
	usr1 := Usuario{
		nombre: "Boris",
		email:  "BorisMR001@gmail.com",
		activo: true,
	}

	usr2 := Usuario{
		nombre: "Doris",
		email:  "DorisRM002@gmail.com",
		activo: true,
	}

	etd1 := Estudiante{
		usuario: usr1,
		codigo:  "00001a",
	}

	etd2 := Estudiante{
		usuario: usr2,
		codigo:  "00002e",
	}

	fmt.Println(etd1, etd2)

	// probando relacion de 1 a *
	video1 := Video{titulo: "Introduccion"}
	video2 := Video{titulo: "Desarrollo"}

	curso := Curso{
		titulo: "Probando Golang",
		videos: []Video{video1, video2},
	}

	video1.curso = curso
	video2.curso = curso

	fmt.Println("Curso:", curso.titulo)

	for _, video := range curso.videos {
		fmt.Println(video.titulo)
	}
}
