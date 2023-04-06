package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	// Http Client instanciado
	c := http.Client{}

	// Cria o JSON para o POST
	jsonVar := bytes.NewBuffer([]byte(`{"name": "pery"}`))
	resp, err := c.Post("http://google.com", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Joga para a tela a resposta
	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
