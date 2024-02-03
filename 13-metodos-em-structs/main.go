package main

import (
	"fmt"
)

type Endereco struct {
	Logradoura string
	Numero     int
	Cidade     string
}

type Client struct {
	Nome     string
	Idade    int
	Ativo    bool
	Endereco          // composição
	Address  Endereco // nova propriedade
}

func (c Client) Desativar() Client {
	c.Ativo = false
	return c
}

func main() {
	renan := Client{Ativo: true, Idade: 30, Nome: "Renan"}

	fmt.Printf("nome: %s, idade: %d, ativo: %t\n", renan.Nome, renan.Idade, renan.Ativo)
	renan.Idade = 35
	fmt.Printf("nova idade do renan: %d\n", renan.Idade)
	renan.Cidade = "São Paulo"
	renan.Address.Cidade = "São Paulo------"
	fmt.Printf("cidade da composição: %s\n", renan.Cidade)
	fmt.Printf("cidade da propriedade: %s\n", renan.Address.Cidade)

	renan = renan.Desativar()
	fmt.Printf("Renan está: %t\n", renan.Ativo)

}
