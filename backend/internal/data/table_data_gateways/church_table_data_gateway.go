package tabledatagateways

import (
	"church-for-all/internal/models"
	"database/sql"
	"strings"

	_ "github.com/go-sql-driver/mysql" // Driver de conexão com o banco de dados MySQL
)

// ChurchTableDataGateway é uma estrutura que representa a tabela church no banco de dados
type ChurchTableDataGateway struct {
	db *sql.DB
}

// NewChurchTableDataGateway é uma função que retorna uma instância da estrutura ChurchTableDataGateway
func NewChurchTableDataGateway() *ChurchTableDataGateway {
	// Conexão com o banco de dados MySQL
	dsn := "root:root@tcp(localhost:3306)/church_for_all?parseTime=true"
	db, err := sql.Open("mysql", dsn)

	// Responsável por verificar se ocorreu algum erro ao conectar no banco de dados
	if err != nil {
		panic("Erro ao conectar no banco de dados: " + err.Error())
	}

	// Responsável por retornar uma instância da estrutura ChurchTableDataGateway
	return &ChurchTableDataGateway{db: db}
}

// CreateChurch é uma função que cria uma igreja
func (c *ChurchTableDataGateway) CreateChurch(church models.Church) error {
	phonesStr := strings.Join(church.Phones, ",")
	_, err := c.db.Exec("INSERT INTO church (id, name, address, phones, email) VALUES (?, ?, ?, ?, ?)", church.ID, church.Name, church.Address, phonesStr, church.Email)
	return err
}

// GetChurchs é uma função que busca todas as igrejas
func (c *ChurchTableDataGateway) GetChurchs() ([]models.Church, error) {
	rows, err := c.db.Query("SELECT * FROM church")
	if err != nil {
		return nil, err
	}

	var churches []models.Church
	for rows.Next() {
		var church models.Church
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
