package main
//Usando uma literal composta:
// - Crie um array que suporte 5 valores do tipo int
// - Atribua valores aos seus índices
//Utilize range e demonstre os valores do array
//Utilizando format printing, demonstre o tipo do array
import "fmt"
 
func main(){
	array := [5]int{1, 2, 3, 4, 5}

	for i, v := range array{
		fmt.Println(i, v)		
	}
	fmt.Printf("%T", array)
}