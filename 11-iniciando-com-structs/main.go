package main

import (
	"fmt"
)

type Client struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	renan := Client{Ativo: true, Idade: 30, Nome: "Renan"}

	fmt.Printf("nome: %s, idade: %d, ativo: %t\n", renan.Nome, renan.Idade, renan.Ativo)
	renan.Idade = 35
	fmt.Printf("nova idade do renan: %d\n", renan.Idade)

}
