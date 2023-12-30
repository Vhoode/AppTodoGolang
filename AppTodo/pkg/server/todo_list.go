package server

import (
	reapeatTodo "repeat/Todo"
	"repeat/Todo/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list reapeatTodo.TodoList) (int, error) {
	return s.repo.Create(int(userId), list)
}
