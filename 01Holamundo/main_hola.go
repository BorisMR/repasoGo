package main

//import de librerias, ya sean bases del lenguaje y sistema o externas de github por ejemplo
import "fmt"

//funcion principal, solo se permite una funcion main por programa
func main() {
	var variable string = "ASDF"
	var variable_entera int = 100

	fmt.Println("Hola mundo con salto de linea")

	//pruebas usando libreria fmt
	fmt.Printf("Texo con formato %s y entero %d", variable, variable_entera)

	var msj string = fmt.Sprintf("\nTexo con formato %s y entero %d guardado en variable", variable, variable_entera)
	fmt.Println(msj)

	fmt.Printf("Tipo %T , Valor %d", variable_entera, variable_entera)

	fmt.Print("Probando leer de consola, ingrese texto: ")
	fmt.Scanln(&variable)
	fmt.Println("Modificado el texto ASDF, ahora es:", variable)
}
