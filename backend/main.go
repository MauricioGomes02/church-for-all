package main

import (
	"encoding/json"
	"net/http"
)

// Church representa uma igreja
type Church struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Address string   `json:"address"`
	Phones  []string `json:"phones"`
	Email   string   `json:"email"`
}

func main() {
	// Responsável por definir o comportamento da aplicação ao acessar a url localhost:8080/churchs
	http.HandleFunc("/churchs", func(w http.ResponseWriter, r *http.Request) {
		// Responsável por definir o comportamento da aplicação quando o método POST é utilizado
		if r.Method == http.MethodPost {
			// Responsável por definir o tipo de conteúdo da resposta
			w.Header().Set("Content-Type", "application/json")

			// Responsável por decodificar o corpo da requisição e armazenar em uma variável
			var church Church
			err := json.NewDecoder(r.Body).Decode(&church)

			// Responsável por verificar se ocorreu algum erro ao decodificar o corpo da requisição
			if err != nil {
				// Responsável por retornar um erro 400 caso o corpo da requisição seja inválido
				http.Error(w, "Invalid request body", http.StatusBadRequest)
				return
			}

			// Responsável por retornar a igreja cadastrada
			json.NewEncoder(w).Encode(church)
		}
	})

	// Responsável por iniciar o servidor e escutar na porta 8080
	http.ListenAndServe(":8080", nil)
}
