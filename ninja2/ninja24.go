package main 
//Crie um programa que:
// - Atribua um valor int a uma variável
// - Demonstre este valor em decimal, binário, e hexadecimal
// - Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
// - Demonstre esta outra variável em decimal, binário e hexadecimal

import "fmt"

var r int = 25

func main() {
	fmt.Printf("%d\t%b\t%#x\n", r, r, r)
	
	a := r << 1
	fmt.Printf("%d\t%b\t%#x", a, a, a)
}