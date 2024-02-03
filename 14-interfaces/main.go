package main

import (
	"fmt"
)

type Endereco struct {
	Logradoura string
	Numero     int
	Cidade     string
}

type Pessoa interface {
	Desativar()
}

type Client struct {
	Nome     string
	Idade    int
	Ativo    bool
	Endereco          // composição
	Address  Endereco // nova propriedade
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {
	fmt.Print("Não estou fazendo nada")
}

func (c Client) Desativar() {
	c.Ativo = false
	fmt.Printf("O Cliente %s foi desativado.\n", c.Nome)
	//return c
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
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

	//renan = renan.Desativar()
	minhaEmpresa := Empresa{}
	Desativacao(renan)
	Desativacao(minhaEmpresa)

	//fmt.Printf("Renan está: %t\n", renan.Ativo)

}
