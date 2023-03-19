package main
//Crie um programa que utilize a declarção switch, sem switch statement especificado

import "fmt"

func main(){
	x := "sábadou"

	switch {
		case x == "segunda":
			fmt.Print("Dia útil")
		case x == "terça":
			fmt.Print("Dia útil")
		case x == "quarta":
			fmt.Print("Dia útil")
		case x == "quinta":
			fmt.Print("Dia útil")
		case x == "sexta":
			fmt.Print("Dia útil")
		case x == "sábado":
			fmt.Print("Final de semana")
		case x == "domingo":
			fmt.Print("Final de semana")
		default:
			fmt.Print("Dia da semana inválido!")
	}

}