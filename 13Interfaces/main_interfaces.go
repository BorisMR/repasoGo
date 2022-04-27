package main

import "fmt"

type Perro struct{}
type Pez struct{}
type Pajaro struct{}

// maneja metodos similares para un
type Animal interface {
	mover() string
}

func (*Perro) mover() string {
	return "Perro caminando"
}

func (*Pez) mover() string {
	return "Pez nadando"
}

func (*Pajaro) mover() string {
	return "Pajaro volando"
}

func moverAnimal(animal Animal) {
	fmt.Println(animal.mover())
}

func main() {
	perrito := Perro{}
	moverAnimal(&perrito)
	pececillo := Pez{}
	moverAnimal(&pececillo)
	pajarito := Pajaro{}
	moverAnimal(&pajarito)
}
