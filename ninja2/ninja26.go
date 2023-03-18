package main 
//Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
//Demonstre esses valores
import "fmt"

const (
	year = 2023 + iota
	year2
	year3 
	year4 
)

func main(){
	fmt.Println(year)
	fmt.Println(year2)
	fmt.Println(year3)
	fmt.Println(year4)

	x := 31337
	fmt.Printf("%#x", x)
}