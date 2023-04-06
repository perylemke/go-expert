package main

import (
	"io"
	"net/http"
)

func main() {
	// Http Client instanciado
	c := http.Client{}
	req, err := http.NewRequest("GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}

	// Pode setar um header específico para a requisição
	req.Header.Set("Accept", "application/json")

	// Realiza a requisção
	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Exibe na tela
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
