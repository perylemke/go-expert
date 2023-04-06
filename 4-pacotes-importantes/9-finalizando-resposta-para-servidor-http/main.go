package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type ViaCEP struct {
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
	// Rota da API
	http.HandleFunc("/", BuscaCepHandler)

	// Sobe o servidor HTTP
	http.ListenAndServe(":8080", nil)
}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	// Verifica se é colocado qualquer coisa e retorna 404
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Pega o parametro ?cep= e verifica se está preenchido, se não estiver retorna 400
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Chama a funcao BuscaCep para consultar o CEP pela API.
	cep, err := BuscaCep(cepParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Setando cabeçalho
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// res, err := json.Marshal(cep)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// w.Write(res)

	// A função abaixo tem o mesmo intuito a comentada acima.
	json.NewEncoder(w).Encode(cep)
}

func BuscaCep(cep string) (*ViaCEP, error) {
	// Realiza a consulta no ViaCEP e usa o defer para evitar vazamentos de recursos
	resp, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Faz a leitura da respose.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Realiza o unmarshal da response e retorna o ponteiro da struct
	var c ViaCEP
	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
