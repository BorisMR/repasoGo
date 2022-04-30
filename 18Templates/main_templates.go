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

// ParseGlob nbos sirve para cargar multiples archivos a traves de una ruta
var templates = template.Must(template.New("T").ParseGlob("plantillas/*.html"))

func renderizarTemplate(rw http.ResponseWriter, name string, data interface{}) {
	err := templates.ExecuteTemplate(rw, name, data)

	if err != nil {
		http.Error(rw, "No se pudo renderizar", http.StatusInternalServerError)
	}
}

func Index(rw http.ResponseWriter, rq *http.Request) {
	// creamos una data
	usrData := Usuarios{
		UserName:    "Boris",
		Edad:        31,
		TieneSuenno: false,
	}

	renderizarTemplate(rw, "index.html", usrData)
}

func OtraPagina(rw http.ResponseWriter, rq *http.Request) {
	renderizarTemplate(rw, "otrapagina.html", nil)
}

func main() {
	// mux
	mux := http.NewServeMux()

	// ruta
	mux.HandleFunc("/", Index)
	mux.HandleFunc("/otrapagina", OtraPagina)

	server := &http.Server{
		Addr:    "localhost:1337",
		Handler: mux,
	}

	// Iniciando Servidor
	fmt.Println("Servidor para template iniciado en puerto 1337")
	log.Fatal(server.ListenAndServe())
}
