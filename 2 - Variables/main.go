package main

import "fmt"

func main() {
	var variavel1 string = "Inferencia explicita"
	variavel2 := "Inferencia implicita"

	fmt.Println(variavel1)
	fmt.Println(variavel2)
	
	var (
		variavel3 string = "Variavel 3"
		variavel4 string = "Variavel 4"
	)

	fmt.Println(variavel3)
	fmt.Println(variavel4)

	variavel5, variavel6 := "Variavel 5", "Variavel6"

	fmt.Println(variavel5)
	fmt.Println(variavel6)

	const constants1 string = "Mesma regra da variaveis"
	fmt.Println(constants1)

	variavel5, variavel6 = variavel6, variavel5

	fmt.Println(variavel5, variavel6)

	// nao existe char em go, ele imprimi o numero da tabela asc
	char := 'A'
	fmt.Println(char)
}