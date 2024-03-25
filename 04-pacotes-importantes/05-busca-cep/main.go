package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {
		req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Erro ao fazer a req: %v\n", err)
		}

		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
		}

		var data ViaCep
		err = json.Unmarshal(res, &data)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error parsing: %v\n", err)
		}

		fmt.Println(data)

		file, err := os.Create("cidade.txt")
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
		}
		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s", data.Cep, data.Localidade))
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error writing to file: %v\n", err)
		}
	}

}
