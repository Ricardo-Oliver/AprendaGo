package main
//Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
// - Nome, - Sobrenome, - Sabores favoritos de sorvete
//Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice
// que contem os sabores de sorvete.

import "fmt"

type pessoa struct {
	nome 			string 
	sobrenome 		string
	sabores			[]string
}

func main() {

	pessoa1 := pessoa{
		nome: "João",
		sobrenome: "Silva",
		sabores: []string{"napolitano", "creme"},
	}
	
	// Usando uma forma reduzida para preencher os dados da struct
	pessoa2 := pessoa{
		"José",
		"Pereira",
		[]string{"chocolate", "morango"},
	}

	fmt.Println("Meu nome é", pessoa1.nome, pessoa1.sobrenome, "e meus sabores favoritos de sorvete são:")
	for _, v := range pessoa1.sabores{
		fmt.Println("-", v)
	}

	fmt.Printf("\n")

	fmt.Println("Meu nome é", pessoa2.nome, pessoa2.sobrenome, "e meus sabores favoritos de sorvete são:")
	for _, v := range pessoa2.sabores{
		fmt.Println("-", v)
	}
}