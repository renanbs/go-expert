package main

import "fmt"

func main() {
	var minhaVar interface{} = "Renan BS"
	println(minhaVar)
	println(minhaVar.(string))

	res, ok := minhaVar.(int)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v\n", res, ok)

	res2 := minhaVar.(int)
	fmt.Printf("O valor de res2 é %v, mas vai dar panic\n", res2)
}
