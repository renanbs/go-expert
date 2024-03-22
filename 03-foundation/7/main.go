package main

import "fmt"

func main() {
	salarios := map[string]int{"Wesley": 100, "Juca": 200}
	fmt.Println(salarios["Wesley"])
	delete(salarios, "Wesley")
	salarios["Wes"] = 5000
	fmt.Println(salarios["Wesley"])
	fmt.Println(salarios["Wes"])

	//sal := make(map[string]int)
	//sal1 := map[string]int

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salario é %d\n", salario)
	}
}
