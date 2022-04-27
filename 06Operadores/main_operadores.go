package main

import "fmt"

func main() {
	/*
		relacionales == != > >= < <=
		logicos && ||
	*/

	var num1 int = 66
	var num2 int = 63

	//ejemplo de if con anidado
	if num1 > num2 {
		fmt.Println("num1:", num1, "es mayor")
	} else if num1 == num2 {
		fmt.Println("num1 y num2:", num1, "son iguales")
	} else {
		fmt.Println("num2:", num2, "es mayor")
	}

	//ejemplo de if con declaracion en condicion con scope limitado al if
	if num3 := 100; num3 > num1 && num3 > num2 {
		fmt.Println("num3:", num3, "es mayor a num1 y num2")
	}

	var letra string = "a"

	// probando un switch
	switch letra {
	case "a", "A", "e", "E", "i", "I", "o", "O", "u", "U":
		fmt.Println(letra, "es vocal")
	default:
		fmt.Println(letra, "no es vocal")
	}
}
