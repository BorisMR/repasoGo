package main

import (
	"fmt"
)

func main() {
	//inicializacion con asignacion posterior
	var nombre string
	nombre = "Sr.Boris"

	//inicializacion con asignacion inmediata
	var apellido string = "Morales"

	//inicializacion abreviada e inferida
	edad := 31

	//inicializacion de multiples tipos, ojo en no existe undefined
	var cadena string = "cadena"
	var entero int = 100
	var flotante_64 float64 = 10.123321
	var boleana bool = true

	const clasico_pi = 3.141592

	//test -> cadena 100 10.123321 1
	fmt.Println(cadena, entero, flotante_64, boleana)
	//test -> Sr.Boris Morales 31
	fmt.Println(nombre, apellido, edad)
	//test -> 3.141592
	fmt.Println(clasico_pi)

	//- - - - ARITMETICOS
	a := 3
	b := 2
	result_suma := a + b
	result_resta := a - b
	result_multi := a * b
	var result_divi float32
	result_divi = 3.0 / 2.0
	result_mod := a % b

	fmt.Println("A:", a, "B:", b)
	fmt.Println("Suma:", result_suma)
	fmt.Println("Resta:", result_resta)
	fmt.Println("Multiplicacion:", result_multi)
	fmt.Println("Division:", result_divi)
	fmt.Println("Modulo:", result_mod)
}
