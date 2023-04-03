package main

import "net/http"

func main() {
	// Geralmente você pode colocar uma função na chamada.
	http.HandleFunc("/", BuscaCEP)

	// Pode ser criada com uma função anonima
	http.HandleFunc("/anom", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	http.ListenAndServe(":8080", nil)
}

func BuscaCEP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
