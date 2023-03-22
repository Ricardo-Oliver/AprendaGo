package main
//Crie um map com key tipo string e value tipo []string.
// - Key deve conter nomes no formato sobrenome_nome
// - Value deve conter os hobbies favoritos da pessoa
//Demonstre todos esses valores e seus índices

import "fmt"

func main() {

	friends := map[string][]string {
		"Oliveira_Ricardo": {"ouvir músicas","jogos on-line"},
		"Antonio_Joao": {"jogos on-oline", "ouvir música"},
		"Ramos_Victor": {"ex", "ouvir música"},
	}

	for k, v := range friends{
		fmt.Println(k)
		for i, h := range v {
			fmt.Println("\t", i, " - ", h)
		}
	}
	
}