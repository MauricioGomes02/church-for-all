package mysql

import (
	"church-for-all/internal/domain/entities"
	"database/sql"
	"strings"

	_ "github.com/go-sql-driver/mysql" // Driver de conexão com o banco de dados MySQL
)

// CreateChurchMySQL é responsável por encapsular o gerenciamento de dados relativo ao ato de criar uma igreja
type CreateChurchMySQL struct {
	database *sql.DB
}

// NewCreateChurchMySQL é um método responsável por criar uma nova instância de CreateChurchMySQL
func NewCreateChurchMySQL(d *sql.DB) *CreateChurchMySQL {
	return &CreateChurchMySQL{database: d}
}

// Create é um método responsável por salvar o registro de uma igreja no banco de dados
func (c *CreateChurchMySQL) Create(church *entities.Church) error {
	phonesStr := strings.Join(church.Phones, ",")
	_, err := c.database.Exec("INSERT INTO church (id, name, address, phones, email) VALUES (?, ?, ?, ?, ?)", church.ID, church.Name, church.Address, phonesStr, church.Email)
	return err
}
