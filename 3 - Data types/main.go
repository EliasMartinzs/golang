package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = 10000000000
	fmt.Println(numero)

	// uint numero sem sinais
	var numero2 uint32 = 1000
	fmt.Println(numero2)

	// alias
	// INT32 = rune
	var numero3 rune = 0000
	fmt.Println(numero3)

	// byte = uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.12
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 22123.12
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.12
	fmt.Println(numeroReal3)

	var boolean bool = false
	fmt.Println(boolean)

	var erro error
	fmt.Println(erro)

	// error pesonalizado usar um pacote `errors`
	var erro2 error = errors.New("Erro interno")
	fmt.Println(erro2)
}