package main

import saludo "saludador/saludador"

/*
	Se genera archivo go.mod con nombre del paquete
	y version de Go, luego se importa usando
		=> archivo nombreDePaquete/carpeta
		=> nombreDePaquete/carpetaOarchivo <- si son iguales
*/
func main() {
	saludo.DecirHola()
	saludo.Saludar()
	saludo.SaludarViaPrivado()
}
