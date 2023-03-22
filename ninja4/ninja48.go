package main
//Utilizando o exercício anterior, 
//adicione uma entrada ao map e demonstre o map inteiro
//utilizando range

import "fmt"

func main() {
	
	friends := map[string] []string {
		"Oliveira_Ricardo": {"Ouvir música eletrônica", "Jogos on-line"},
		"Ramos_Victor": {"Ex", "Ouvir música", "Jogos on-line"},
		"Antonio_Joao": {"Ouvir música viking", "Jogos on-line"},
	}

	friends["Miranda_Daniel"] = []string{"Jogos on-line", "Pilotar guincho", "Patinar"}

	for k, v := range friends{
		fmt.Println(k)
		for i, h := range v {
			fmt.Println("\t", i, " - ", h)
		}
	}
}