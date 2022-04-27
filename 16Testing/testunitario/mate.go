package testunitario

func Suma(a, b int) int {
	return a + b
}

func GetMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// correr go test en terminal
// correr go test -cover nos muestra el % de covertura de las pruebas
