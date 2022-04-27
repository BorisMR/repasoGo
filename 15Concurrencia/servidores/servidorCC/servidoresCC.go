package main

import (
	"fmt"
	"net/http"
	"time"
)

func revisarServidorConcurrente(servidor string, canal chan string) {
	_, err := http.Get(servidor)

	if err != nil {
		canal <- servidor + " :No disponible"
	} else {
		canal <- servidor + " :Disponible"
	}
}

func main() {
	inicio := time.Now()

	//creamos canales para manejar las Goroutines y la concurrencia, asi aseguramos su comunicacion y sincronizacion
	canal := make(chan string)

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
		// go se utiliza para generar Goroutines y poder ejecutar concurrentemente
		go revisarServidorConcurrente(servidor, canal)
	}

	for i := 0; i < len(servidores); i++ {
		fmt.Println(<-canal)
	}

	finTiempo := time.Since(inicio)
	fmt.Println("Tiempo verificando:", finTiempo)
}
