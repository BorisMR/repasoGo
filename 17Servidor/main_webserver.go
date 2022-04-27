package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handlers <- metodos para manejar las peticiones
func Hola(rw http.ResponseWriter, rq *http.Request) {
	fmt.Println(rq.Method, rq.RequestURI)
	fmt.Fprintln(rw, "Hola mundo")
}

func Prueba(rw http.ResponseWriter, rq *http.Request) {
	fmt.Println(rq.Method, rq.RequestURI)
	fmt.Fprintln(rw, "pagina de prueba")
}

func NoEncontrado(rw http.ResponseWriter, rq *http.Request) {
	fmt.Println(rq.Method, rq.RequestURI)
	http.NotFound(rw, rq)
}

func Error(rw http.ResponseWriter, rq *http.Request) {
	fmt.Println(rq.Method, rq.RequestURI)
	http.Error(rw, "Este es un error", http.StatusConflict)
}

func Datos(rw http.ResponseWriter, rq *http.Request) {
	fmt.Println(rq.Method, rq.RequestURI)

	nombre := rq.URL.Query().Get("nombre")
	edad := rq.URL.Query().Get("edad")

	// http://localhost:1337/datos?nombre=boris&edad=31
	fmt.Fprintf(rw, "Hola %s tienes %s años", nombre, edad)
}

func main() {
	// Mux
	mux := http.NewServeMux()

	// Router
	mux.HandleFunc("/", Hola)
	mux.HandleFunc("/prueba", Prueba)
	mux.HandleFunc("/notfound", NoEncontrado)
	mux.HandleFunc("/error", Error)
	mux.HandleFunc("/datos", Datos)

	server := &http.Server{
		Addr:    "localhost:1337",
		Handler: mux,
	}

	// Iniciando Servidor
	fmt.Println("Servidor iniciado en puerto 1337")
	log.Fatal(server.ListenAndServe())
}

/*
Reiniciar servidor automaticamnete


*/
