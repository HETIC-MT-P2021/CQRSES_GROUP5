package users

import "database/sql"

//RepositoryInterface interface for functions for user repository
type RepositoryInterface interface {
}

//Repository struct for db connection
type Repository struct {
	Conn *sql.DB
}

//Close the DB connection
func (repository *Repository) Close() {
	repository.Conn.Close()
}
