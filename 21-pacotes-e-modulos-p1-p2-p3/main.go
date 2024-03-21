package main

import (
	"fmt"
	"github.com/renanbs/go-expert/21-pacotes-e-modulos-p1-p2-p3/matematica"
)

func main() {
	soma := matematica.Soma(1, 2)
	fmt.Printf("O resultado: %v\n", soma)
	fmt.Println(matematica.A)
}
