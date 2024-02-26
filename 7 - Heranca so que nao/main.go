package main

import "fmt"

type pessoa struct {
	nome     string
	idade    uint8
	endereco string
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	p1 := pessoa{"Golias", 22, "Rua tal"}

	el := estudante{p1, "ads", "uninter"}
	fmt.Println(el)
}