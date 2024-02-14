package main

import "fmt"

type Cliente struct {
	nome string
}

func (c Cliente) andou() {
	c.nome = "Renan BS"
	fmt.Printf("O Cliente %v andou\n", c.nome)
}

type Conta struct {
	saldo int
}

func NewConta() *Conta {
	return &Conta{saldo: 0}
}

func (c Conta) simular(valor int) int {
	c.saldo += valor
	return c.saldo
}

func (c *Conta) simular2(valor int) int {
	c.saldo += valor
	return c.saldo
}

func main() {
	renan := Cliente{
		nome: "Renan",
	}

	renan.andou()
	fmt.Printf("O Valor da struct com nome é %v\n", renan.nome)

	conta := Conta{saldo: 100}
	conta.simular(200)
	println(conta.saldo)
	conta.simular2(200)
	println(conta.saldo)

	conta2 := NewConta()
	conta2.simular2(10)
	println(conta2.saldo)
}
