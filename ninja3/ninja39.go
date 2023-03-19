package main
//Crie um programa que utilize a declaração switch, 
//onde o switch statement seja uma variável do tipo string e identificador "esportefavorito"

import "fmt"

func main(){
	esportefavorito := "natação"
	switch esportefavorito {
	case "luta":
		fmt.Print("Não é meu esporte favorito")
	case "ciclismo":
		fmt.Print("Não é meu esporte favorito")
	case "natação":
		fmt.Print("Meu esporte favorito")
	default:
		fmt.Print("Nunca vi esse esporte :(")
	}
}