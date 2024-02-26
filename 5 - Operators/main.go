package main

import "fmt"

func main() {
	// Aritmeticos
	soma := 1 + 1
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma,subtracao,divisao,multiplicacao,restoDaDivisao)

	// Atribuicao
	var variavel1 string = "STRING"
	variavel2 := "STRING"
	fmt.Println(variavel1,variavel2)

	// Operadores Relacionais
	fmt.Println(1>2)
	fmt.Println(1>=2)
	fmt.Println(1==2)
	fmt.Println(1<=2)
	fmt.Println(1<2)
	fmt.Println(1!=2)

	//  Operadores logicos
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!falso)

	// Operadores Unarios
	numero := 10
	numero++
	numero--
	numero += 16
	fmt.Println(numero)

	
}