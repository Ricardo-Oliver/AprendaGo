package main
//Crie um slice contendo slices de strings ([][]string).
//Atribua valores a este slice multi-dimensional da seguinte maneira:
//"Nome", "Sobrenome", "Hobby favorito"
//Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.

import "fmt"

func main() {
	peoples := [][]string{
		{"Ricardo", "Oliveira", "Música"},
		{"João", "Antonio", "Jogos on-line"},
		{"Victor", "Ramos", "Ex"},
	}

	fmt.Println(peoples)
	
	for _, v := range peoples{
		fmt.Println(v[0])
		for _, item := range v{
			fmt.Println("\t", item)
		}
	}
}