package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		// recover para controlar "panic" y controlar cierres forzosos
		if error := recover(); error != nil {
			fmt.Println("Error al finalizar el programa")
		}
	}()

	if archivo, error := os.Open("archivo.txt"); error != nil {
		// muestra error y detiener ejecucion de programa
		panic("No se encontro archivo")
	} else {
		fmt.Println("Leyendo archivo ...")

		// funcion anonima con defer para cerrar archivo al terminar lectura
		defer func() {
			fmt.Println("Cerrando archivo ...")
			archivo.Close()
		}()

		// contenido de archivo esta obtenido en bytes
		contenido := make([]byte, 254)
		largo, _ := archivo.Read(contenido)

		//convertimos contenido a strings para mostrarlo
		contenidoArchivo := string(contenido[:largo])
		fmt.Println(contenidoArchivo)
	}
}
