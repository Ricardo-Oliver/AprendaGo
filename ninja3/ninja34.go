package main
//Crie um loop utilizando a sixtaxe: for {}
//Utilize-o para demonstrar os anos desde que você nasceu
import "fmt"

func main(){
	anoNascimento := 1997
	anoAtual := 2023

	for {
		if anoNascimento > anoAtual {
			break
		}
		fmt.Println(anoNascimento)
		anoNascimento++
	}
}