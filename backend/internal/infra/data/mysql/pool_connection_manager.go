package mysql

import "database/sql"

// PoolConnectionManager é responsável por encapsular a conexão com o pool e fornecer o acesso a manipulação dos dados
type PoolConnectionManager struct {
	Database *sql.DB
}

// NewPoolConnectionManager é um método responsável por criar uma nova instância de PoolConnectionManager
func NewPoolConnectionManager() *PoolConnectionManager {
	// Conexão com o banco de dados MySQL
	dsn := "root:root@tcp(localhost:3306)/church_for_all?parseTime=true"
	db, err := sql.Open("mysql", dsn)

	// Responsável por verificar se ocorreu algum erro ao conectar no banco de dados
	if err != nil {
		panic("Erro ao conectar no banco de dados: " + err.Error())
	}

	return &PoolConnectionManager{Database: db}
}
