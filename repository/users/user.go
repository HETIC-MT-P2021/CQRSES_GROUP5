package users

import (
	"database/sql"
	"github.com/HETIC-MT-P2021/gocqrs/models"
)

//GetUser is for getting a user by username
func (repository *Repository) GetUser(username string) (*models.User, error) {
	row := repository.Conn.QueryRow(`
		SELECT 
		u.id, 
		u.username, 
		u.email, 
		u.password, 
		u.role 
		FROM api_user u 
		WHERE u.username=(?)`, username)
	var user models.User
	switch err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role); err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return &user, nil
	default:
		return nil, err
	}
}

//SaveUser is for saving a new user
func (repository *Repository) SaveUser(user *models.User) error {
	stmt, err := repository.Conn.Prepare(`
	INSERT INTO api_user 
	(username, email, password, role) 
	VALUES(?,?,?,?) `)
	if err != nil {
		return err
	}

	res, errExec := stmt.Exec(user.Username, user.Email, user.Password, user.Role)
	if errExec != nil {
		return errExec
	}

	lastInsertedID, errInsert := res.LastInsertId()
	if errInsert != nil {
		return errInsert
	}

	user.ID = lastInsertedID

	return nil
}
