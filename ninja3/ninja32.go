package main
//Põe na tela: O unicode code point de todas as letras maiúsculas do alfabeto, três vezes cada.
//Por exemplo: 65 - U+0041 'A' 
//					U+0041 'A'
//					U+0041 'A'  
import "fmt"

func main() {
	for x := 65; x <= 90; x++{
		fmt.Println(x)
		for y := 1; y <= 3; y++{
			fmt.Printf("\t%#U\n", x)
		}
	}
}