package main

import "fmt"

func somar(n1 int, n2 int) int {
	return n1 + n2
}

func doisRetornos (n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	sub := n1-n2

	return soma,sub
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

// resultadoSoma, resultadoSubtracao := doisRetornos(10,20)
_, resultadoSubtacao := doisRetornos(10,20)
fmt.Println(resultadoSubtacao)
}