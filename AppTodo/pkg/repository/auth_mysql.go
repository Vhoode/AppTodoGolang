package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	reapeatTodo "repeat/Todo"
)

type AuthMysql struct {
	db *sql.DB
}

func (r *AuthMysql) GetUser(username, password string) (reapeatTodo.User, error) {
	var user reapeatTodo.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE username = ? AND password_hash = ?", usersTable)

	stmt := r.db.QueryRow(query, username, password)
	err := stmt.Scan(&user.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logrus.Errorf("user not found in db:%s", err.Error())
			return user, err
		}
	}
	return user, err
}

func NewAuthMysql(db *sql.DB) *AuthMysql {
	return &AuthMysql{db: db}
}

func (r *AuthMysql) CreateUser(user reapeatTodo.User) (int64, error) {
	var id int64

	query := "INSERT INTO user_list (name, username, password_hash) VALUES (?, ?, ?)"
	result, err := r.db.Exec(query, user.Name, user.UserName, user.Password)
	if err != nil {
		return 0, err
	}

	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
