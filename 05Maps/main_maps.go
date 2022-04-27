package main

import "fmt"

func main() {
	// los mapas son colecciones infinitas de elementos identificados por claves
	meses := make(map[int]string)
	dias := make(map[string]int)

	meses[-1] = "enero"
	meses[2] = "febrero"
	meses[3] = "marzo"
	meses[4] = "abril"
	meses[5] = "mayo"

	dias["lunes"] = 0
	dias["martes"] = 1

	// map[-1:enero 2:febrero 3:marzo 4:abril 5:mayo]
	fmt.Println(meses)

	//borrando elemento (map, key)
	delete(meses, -1)

	// map[2:febrero 3:marzo 4:abril 5:mayo] 4
	fmt.Println(meses, len(meses))
	// map[lunes:0 martes:1]
	fmt.Println(dias)
}
