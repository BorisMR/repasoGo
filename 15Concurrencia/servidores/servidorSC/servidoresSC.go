package main

import (
	"fmt"
	"net/http"
	"time"
)

func revisarServidorSinConcurrencia(servidor string) {
	_, err := http.Get(servidor)

	if err != nil {
		fmt.Println("Servidor no disponible")
	} else {
		fmt.Println("Servidor funcionando", servidor)
	}
}

func main() {
	inicio := time.Now()

	servidores := []string{
		"https://www.google.cl",
		"https://www.youtube.com",
		"https://www.facebook.com",
		"https://www.udemy.com",
		"https://www.instagram.com",
		"https://www.buscalibre.cl",
	}

	// ejecucion secuencial sin concurrencia
	for _, servidor := range servidores {
		revisarServidorSinConcurrencia(servidor)
	}

	finTiempo := time.Since(inicio)
	fmt.Println("Tiempo verificando:", finTiempo)
}
