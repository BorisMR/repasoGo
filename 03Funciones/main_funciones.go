package main

import "fmt"

func main() {
	var num1 int = 100
	var num2 int = 28
	var result int = 0

	MostrarSaludo()
	MostrarSaludoNombre("Sr.Boris")
	result = SumarDosNumeros(num1, num2)

	fmt.Printf("Suma de %v y %v es igual a %v", num1, num2, result)
}

func MostrarSaludo() {
	fmt.Println("Hola, sin parametros")
}

func MostrarSaludoNombre(nombre string) {
	fmt.Printf("Hola %s", nombre)
}

func SumarDosNumeros(n1, n2 int) int {
	return n1 + n2
}
