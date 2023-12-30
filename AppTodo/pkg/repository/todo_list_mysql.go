package repository

import (
	"database/sql"
	"fmt"
	reapeatTodo "repeat/Todo"
)

type TodoListMysql struct {
	db *sql.DB
}

func NewTodoListMysql(db *sql.DB) *TodoListMysql {
	return &TodoListMysql{db: db}
}

func (r *TodoListMysql) Create(userId int, list reapeatTodo.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES (?, ?)", todoListTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES (?, ?)", usersTableLists)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}
