package main
//Utilizando a soluçaõ anterior, adicione as opções else if e else

import "fmt"

func main(){
	x := 10
	if x % 2 == 0{
		fmt.Println("Número par!")
	} else if x % 2 != 0 {
		fmt.Println("Número ímpar!")
	} else {
		fmt.Println("Número inválido!")
	}
}