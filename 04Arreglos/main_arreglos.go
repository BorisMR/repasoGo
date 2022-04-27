package main

import "fmt"

func main() {
	//al declarar un arreglo sin sus variables, este las predefine en cero
	var arreglo_vacio [4]int
	// arreglo de strings definido de otra forma
	colores := [4]string{
		"rojo",
		"negro",
		"blanco",
		"amarillo"}

	// arreglo sin longitud definida
	consolas := [...]string{
		"nintendo",
		"playstation",
		"dreamcast",
	}

	// arreglo sin longitud definida, pero con valores en sus indices
	frutas := [...]string{
		0: "palta",
		2: "pera",
		3: "uva",
		4: "aji",
		6: "tomate",
	}

	// [0 0 0 0]
	fmt.Println(arreglo_vacio)

	arreglo_vacio[1] = 100
	arreglo_vacio[2] = 66

	// [0 100 66 0]
	fmt.Println(arreglo_vacio)
	// 66
	fmt.Println(arreglo_vacio[2])
	// [rojo negro blanco amarillo]
	fmt.Println(colores)
	// [nintendo playstation dreamcast]
	fmt.Println(consolas)
	// [palta  pera uva] 7 <- largo de arreglo *espacio entre palta y pera
	fmt.Println(frutas, len(frutas))
	// [pera uva] - este imprime un sub arreglo
	fmt.Println(frutas[2:4])
	// [uva aji  tomate] - este imprime desde indice hasta el final
	fmt.Println(frutas[3:])

	/*
		creando slice con funcion make
		-> (tipo, largo, capacidad)
	*/
	numeros := make([]int, 3, 3)

	numeros[0] = 20
	numeros[1] = 30
	numeros[2] = 10

	// [20 30 10] 3 3
	fmt.Println(numeros, len(numeros), cap(numeros))

	// como sobrepasamos tamaño, se usa append y capacidad se duplica automaticamente
	numeros = append(numeros, 66)

	// [20 30 10 66] 4 6
	fmt.Println(numeros, len(numeros), cap(numeros))
}
