package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
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
	// Conexão com o banco de dados MySQL
	dsn := "root:root@tcp(localhost:3306)/church_for_all?parseTime=true"
	db, err := sql.Open("mysql", dsn)

	// Responsável por verificar se ocorreu algum erro ao conectar no banco de dados
	if err != nil {
		panic("Erro ao conectar no banco de dados: " + err.Error())
	}

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

			// Responsável por inserir a igreja no banco de dados
			// Converte o slice de phones em uma string separada por vírgula
			phonesStr := ""
			if len(church.Phones) > 0 {
				phonesStr = strings.Join(church.Phones, ",")
			}
			_, err = db.Exec("INSERT INTO church (id, name, address, phones, email) VALUES (?, ?, ?, ?, ?)", church.ID, church.Name, church.Address, phonesStr, church.Email)
			if err != nil {
				// Responsável por retornar um erro 500 caso ocorra um erro ao inserir a igreja no banco de dados
				http.Error(w, "Failed to insert church", http.StatusInternalServerError)
				return
			}

			// Responsável por retornar a igreja cadastrada
			json.NewEncoder(w).Encode(church)
		}

		// Responsável por definir o comportamento da aplicação quando o método GET é utilizado
		if r.Method == http.MethodGet {
			// Responsável por definir o tipo de conteúdo da resposta
			w.Header().Set("Content-Type", "application/json")

			// Responsável por buscar todas as igrejas no banco de dados
			rows, err := db.Query("SELECT * FROM church")
			if err != nil {
				// Responsável por retornar um erro 500 caso ocorra um erro ao buscar as igrejas no banco de dados
				http.Error(w, "Failed to get churches", http.StatusInternalServerError)
				return
			}

			// Responsável por criar um slice de igrejas
			var churches []Church

			// Responsável por iterar sobre as linhas retornadas pela query
			for rows.Next() {
				var church Church
				var phonesStr string
				err = rows.Scan(&church.ID, &church.Name, &church.Address, &phonesStr, &church.Email)
				if err != nil {
					// Responsável por retornar um erro 500 caso ocorra um erro ao scanear as linhas retornadas pela query
					http.Error(w, "Failed to scan church", http.StatusInternalServerError)
					return
				}

				// Converte a string de telefones separada por vírgula em um slice de strings
				if phonesStr != "" {
					church.Phones = strings.Split(phonesStr, ",")
				} else {
					church.Phones = []string{}
				}
				churches = append(churches, church)
			}

			// Responsável por retornar a lista de igrejas
			json.NewEncoder(w).Encode(churches)
		}
	})

	// Responsável por iniciar o servidor e escutar na porta 8080
	http.ListenAndServe(":8080", nil)
}
