package repository

import (
	"database/sql"
	reapeatTodo "repeat/Todo"
)

type Authorization interface {
	CreateUser(user reapeatTodo.User) (int64, error)
	GetUser(username, password string) (reapeatTodo.User, error)
}

type TodoList interface {
	Create(userId int, list reapeatTodo.TodoList) (int, error)
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthMysql(db),
		TodoList:      NewTodoListMysql(db),
	}
}
