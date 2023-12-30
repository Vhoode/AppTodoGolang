package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	usersTable      = "user_list"
	usersTableLists = "users_lists"
	todoListTable   = "todo_list"
	todoItemTable   = "todo_item"
)

type Config struct {
	Host     string
	Port     uint32
	Username string
	Password string
	DBName   string
}

func NewMySqlDB(cfg Config) (*sql.DB, error) {
	connectionURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	db, err := sql.Open("mysql", connectionURI)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return db, err
	}
	return db, nil
}
