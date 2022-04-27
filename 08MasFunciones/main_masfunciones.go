package main

import (
	"fmt"
	"strings"
)

//funcion variatica que recibe n argumentos usando ...
func Sumar(quien string, numeros ...int) (string, int) {
	mensaje := fmt.Sprintf("Don %s sumo", quien)
	var total int

	// lla variable numeros se opera como si fuese un arreglo de elementos
	for _, numero := range numeros {
		total += numero
	}

	fmt.Println(numeros)

	//retornar ambos valores
	return mensaje, total
}

func Factorial(n int) int {
	if n == 0 {
		return 1
	}

	fact := n * Factorial(n-1)

	return fact
}

func ClosejureRepetir(repeticiones int) func(msj string) string {

	return func(msj string) string {
		return strings.Repeat(msj, repeticiones)
	}
}

func main() {
	//probando funcion variatica
	mostrar, result := Sumar("Boris", 10, 20, 30)
	fmt.Println(mostrar, result)

	//probando factorial recursivo
	numCalc := 6
	fmt.Println("El factorial de", numCalc, "es", Factorial(numCalc))

	//probando funcion anonima autoejecutable
	func() {
		fmt.Println("Esta es una funcion anonima")
	}()

	otraAnonima := func() {
		fmt.Println("Esta es otra funcion anonima")
	}

	otraAnonima() //aqui se llama a segunda funcion anonima

	// probando funcion closejure -> funcion que retorna otra funcion
	mensajeRepetir := ClosejureRepetir(3)
	fmt.Println(mensajeRepetir("UnClosejure "))
}
