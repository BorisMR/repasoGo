package main

import "fmt"

// se define nombre y atributos de struct "objeto"
type Persona struct {
	nombre string
	edad   int
}

// se definen metodos del "objeto", usar punteros par aenlazar struct y metodo
func (prs *Persona) DatosPersona() {
	fmt.Printf("\nNombre: %s - Edad: %v\n", prs.nombre, prs.edad)
}

//funcion para editar de manera implicita
func (prs *Persona) EditNombre(nuevoNombre string) {
	prs.nombre = nuevoNombre
}

// Simulando herencia de "objeto"
type Trabajador struct {
	Persona
	sueldo int
}

func (tbj *Trabajador) DatosTrabajador() {
	fmt.Printf("\nNombre: %s - Edad: %v - Sueldo: %v\n", tbj.nombre, tbj.edad, tbj.sueldo)
}

// estructura para listas de Trabajadores
type Trabajadores struct {
	trabajadores []*Trabajador
}

func (lt *Trabajadores) AddTrabajador(tbj *Trabajador) {
	lt.trabajadores = append(lt.trabajadores, tbj)
}

func (lt *Trabajadores) DelTrabajador(idx int) {
	lt.trabajadores = append(lt.trabajadores[:idx], lt.trabajadores[idx+1:]...)
}

func (lt *Trabajadores) DatosTrabajadores() {
	fmt.Println("\nNomina de Trabajadores")

	for i, tbj := range lt.trabajadores {
		fmt.Printf("\n[%v] - Trabajador: %s", i, tbj.nombre)
	}
}

func main() {
	//creamos objeto de esta forma si conocemos previamente atributos
	p1 := Persona{"Asimov", 72}
	//creamos objeto de esta otra forma si preferimos definir directamente
	p2 := Persona{
		nombre: "Boris",
		edad:   31,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	//editamos annos de p2 de forma explicita
	p2.edad = 100

	fmt.Println(p2)
	p2.DatosPersona()
	//editamos annos de p2
	p2.EditNombre("SrBoris")
	p2.DatosPersona()

	//creamos empleado para heredar
	tbj1 := Trabajador{}
	tbj1.nombre = "Don Guorker"
	tbj1.edad = 35
	tbj1.sueldo = 2000000
	tbj1.DatosTrabajador()

	//probamos lista
	nomina := Trabajadores{}

	nomina.AddTrabajador(&tbj1)
	nomina.DatosTrabajadores()
	nomina.DelTrabajador(0)
	nomina.DatosTrabajadores()
}
