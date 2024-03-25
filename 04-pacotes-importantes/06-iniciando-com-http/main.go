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
	http.HandleFunc("/", BuscaCepHandler)
	http.ListenAndServe(":8080", nil)
}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var resp *ViaCep
	resp, err := BuscaCep(cepParam)
	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//result, err := json.Marshal(resp)
	//if err != nil {
	//	w.WriteHeader(http.StatusFailedDependency)
	//	return
	//}
	//w.Write(result)

	json.NewEncoder(w).Encode(resp)
}

func BuscaCep(cep string) (*ViaCep, error) {
	req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Erro ao fazer a req: %v\n", err)
		return nil, err
	}

	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
		return nil, err
	}

	var data ViaCep
	err = json.Unmarshal(res, &data)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error parsing: %v\n", err)
		return nil, err
	}
	return &data, nil
}
