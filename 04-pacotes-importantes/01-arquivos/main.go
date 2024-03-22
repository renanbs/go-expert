package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}

	tamanho, err := f.Write([]byte("Hello World!"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho %d bytes", tamanho)
	tamanho, err = f.WriteString("Hello World!")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho %d bytes\n", tamanho)

	f.Close()

	// Leitura
	arq, err := os.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arq))
	// Leitura
	arq2, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arq2)
	buffer := make([]byte, 3)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove("file.txt")
	if err != nil {
		panic(err)
	}

}
