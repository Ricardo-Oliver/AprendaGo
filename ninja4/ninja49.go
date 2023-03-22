package main
//Utilizando o exercício anterior, remova uma entrada do map e demonstre
//o map inteiro utilizando range

import "fmt"

func main() {

	friends := map[string][]string {
		"Oliveira_Ricardo": {"Ouvir música eletrônica", "Jogos on-line"},
		"Antonio_Joao": {"Ouvir música viking", "Jogos on-line"},
		"Ramos_Victor": {"Ouvir música", "Jogos on-line"},
		"Miranda_Daniel": {"Patinar", "Pilotar guincho"}, 
	}

	delete(friends, "Oliveira_Ricardo")

	for k, v := range friends{
		fmt.Println(k)
		for i, h := range v {
			fmt.Println("\t", i, " - ", h)
		}
	}
}