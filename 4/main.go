package main

import "fmt"

const a = "Hello"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Renan"
	e float64 = 1.2
	f ID      = 1
)

func main() {
	fmt.Printf("O tipo de e é %T\n", e)
	fmt.Printf("O valor de e é %v\n", e)
	fmt.Printf("O valor de f é %T\n", f)

}
