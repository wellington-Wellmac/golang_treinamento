package variaveis

import "fmt"

func main() {
	var variavel1 string = "Variável 1"
	variavel2 := "Variavel 2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	//outro exemplo:
	var (
		Variavel3 string = "lalala"
		Variavel4 string = "lalala"
	)

	fmt.Println(Variavel3, Variavel4)

	//outro
	variavel5, variavel6 := "Variável 5", "Variavel 6"
	fmt.Println(variavel5, variavel6)
}
