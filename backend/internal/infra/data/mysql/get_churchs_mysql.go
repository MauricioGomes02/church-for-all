package mysql

import (
	"church-for-all/internal/domain/entities"
	"database/sql"
	"strings"

	_ "github.com/go-sql-driver/mysql" // Driver de conexão com o banco de dados MySQL
)

// GetChurchsMySQL é responsável por encapsular o gerenciamento de dados relativo ao ato de obter igrejas
type GetChurchsMySQL struct {
	database *sql.DB
}

// NewGetChurchsMySQL é um método responsável por criar uma nova instância de GetChurchsMySQL
func NewGetChurchsMySQL(d *sql.DB) *GetChurchsMySQL {
	return &GetChurchsMySQL{database: d}
}

// Get é um método responsável por obter os registros de igrejas no banco de dados
func (g *GetChurchsMySQL) Get() ([]entities.Church, error) {
	rows, err := g.database.Query("SELECT * FROM church")
	if err != nil {
		return nil, err
	}

	var churches []entities.Church
	for rows.Next() {
		var church entities.Church
		var phonesStr string
		err = rows.Scan(&church.ID, &church.Name, &church.Address, &phonesStr, &church.Email)
		if err != nil {
			return nil, err
		}
		if phonesStr != "" {
			church.Phones = strings.Split(phonesStr, ",")
		} else {
			church.Phones = []string{}
		}
		churches = append(churches, church)
	}

	return churches, nil
}
