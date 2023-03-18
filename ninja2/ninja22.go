package main
//Escreva expressões utilizando os seguintes operadores, e atribua seus valores a variáveis.
// ==, <=, >=, !=, <, >
//Demonstre estes valores

import "fmt"

func main() {
	a := 10 == 10
	b := 10 <= 9
	c := 10 >= 9
	d := 10 != 10
	e := 10 < 20
	f := 10 > 20

	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n", a, b, c, d, e, f)
}