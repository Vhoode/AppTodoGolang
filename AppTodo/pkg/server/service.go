package server

import (
	reapeatTodo "repeat/Todo"
	"repeat/Todo/pkg/repository"
)

type Authorization interface {
	CreateUser(user reapeatTodo.User) (int64, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type TodoList interface {
	Create(userId int, list reapeatTodo.TodoList) (int, error)
}
type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
