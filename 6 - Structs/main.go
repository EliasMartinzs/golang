package main

import (
	"fmt"
)

type usuario struct {
	nome  string
	idade int
}

func main() {
	var u usuario
	u.nome = "Sei la"
	u.idade = 3

	usuario2 := usuario{"Tendeu", 22}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Struct chama com chaves"}
	fmt.Println(usuario3)

}