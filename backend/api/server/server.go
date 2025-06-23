package server

import (
	"church-for-all/internal/handlers"
	"net/http"
)

// Run é uma função que inicia o servidor
func Run() {
	// Responsável por definir o comportamento da aplicação ao acessar a url localhost:8080/churchs
	http.HandleFunc("/churchs", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handlers.CreateChurch(w, r)
		}

		if r.Method == http.MethodGet {
			handlers.GetChurchs(w, r)
		}
	})

	// Responsável por iniciar o servidor e escutar na porta 8080
	http.ListenAndServe(":8080", nil)
}
