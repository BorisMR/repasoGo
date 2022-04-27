package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var expresion string
	var total int

	fmt.Print("_>")
	fmt.Scanln(&expresion)

	total = Sumar(expresion)

	fmt.Println(total)
}

func Sumar(expresion string) int {
	valores := strings.Split(expresion, "+")
	var resultado int

	for _, valor := range valores {
		// Atoi convierte strings a enteros
		num, err := strconv.Atoi(valor)

		if err != nil {
			fmt.Println(err)
		} else {
			resultado += num
		}
	}

	return resultado
}
