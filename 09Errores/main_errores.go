package main

import (
	"errors"
	"fmt"
)

func Dividir(dividendo int, divisor int) (int, error) {
	if divisor == 0 {
		// manejar error usando libreria errors
		return 0, errors.New("No es posible dividir entre cero")
	} else {
		return dividendo / divisor, nil
	}

}

func main() {
	dividendo := 4
	divisor := 2
	result, error := Dividir(dividendo, divisor)

	if error == nil {
		fmt.Printf("Dividir => %v/%v = %v", dividendo, divisor, result)
	} else {
		fmt.Println(error)
	}
}
