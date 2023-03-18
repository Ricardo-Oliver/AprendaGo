package main
//Escreva um programa que mostre um número em decimal, binário, e hexadecimal

import (
	"fmt"
)

func main() {
	n := 7
	fmt.Printf("%d\t%b\t%#x", n, n, n)
}
