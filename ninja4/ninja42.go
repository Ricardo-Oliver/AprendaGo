package main
//Utilizando como base o exercício anterior, utilize slicing para demonstraros valores:
// - Do primeiro ao terceiro item do slice (incluindo o terceiro item)
// - Do quinto ao último item do slice (incluindo o último item)
// - Do segundo ao sétimo item do slice (incluindo o sétimo item)
// - Do terceiro ao penúltimo item do slice (incluindo o penúltimo item)
// - Desafio: Obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	slice1 := (slice[:3])
	fmt.Println(slice1)

	slice2 := (slice[4:])
	fmt.Println(slice2)

	slice3 := (slice[1:7])
	fmt.Println(slice3)

	slice4 := (slice[2:9])
	fmt.Println(slice4)

	desafio := (slice[2:len(slice) - 1])
	fmt.Println(desafio)
}