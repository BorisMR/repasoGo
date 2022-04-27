package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Usuarios struct {
	UserName    string
	Edad        int
	TieneSuenno bool
}

func Index(rw http.ResponseWriter, rq *http.Request) {
	// creamos un elemento para cargar el archivo
	template, err := template.ParseFiles("templates/index.html")

	// creamos una data
	usrData := Usuarios{
		UserName:    "Boris",
		Edad:        31,
		TieneSuenno: false,
	}

	if err != nil {
		panic(err)
	} else {
		//entregamos la respuesta y adjuntamos la data
		template.Execute(rw, usrData)
	}
}

func main() {
	// mux
	mux := http.NewServeMux()

	// ruta
	mux.HandleFunc("/", Index)

	server := &http.Server{
		Addr:    "localhost:1337",
		Handler: mux,
	}

	// Iniciando Servidor
	fmt.Println("Servidor para template iniciado en puerto 1337")
	log.Fatal(server.ListenAndServe())
}
