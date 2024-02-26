package main

import "fmt"

func main() {
	// Passando uma variavel a outra ele criar uma copia da 1 variavel
	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel2)

	variavel2++

	// Ponteiro e uma referencia de memoria
	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(*ponteiro)

	variavel3 = 200
	fmt.Println(*ponteiro)
}