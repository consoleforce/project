// pkg/store/postgres/postgres.go

package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// ConnectDB устанавливает соединение с базой данных PostgreSQL
func ConnectDB(host, port, user, password, dbname string) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
