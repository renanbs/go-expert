package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)
	if err != nil {
		println(err)
	}
	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		println(err)
	}

	jsonPuro := []byte(`{"Numero": 2, "Saldo": 200}`)
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		println(err)
	}
	println(contaX.Numero)
	println(contaX.Saldo)

	jsonPuro2 := []byte(`{"n": 3, "s": 300}`)
	var contaY Conta
	err = json.Unmarshal(jsonPuro2, &contaY)
	if err != nil {
		println(err)
	}
	fmt.Println(contaY)
	println(contaY.Numero)
	println(contaY.Saldo)

}
