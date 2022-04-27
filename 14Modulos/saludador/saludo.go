package saludo

import "fmt"

const MENSAJE = "Hola desde constante externa"

func DecirHola() {
	fmt.Println("Hola desde modulo saludador, paquete saludo ")
}

func Saludar() {
	fmt.Println(MENSAJE)
}

func SaludarViaPrivado() {
	saludoPrivado()
}

// inicio de nombre con minusculas define scope privado
func saludoPrivado() {
	fmt.Println("Hola desde funcion privada externa saludoPrivado")
}
