package reapeatTodo

type TodoList struct {
	Id          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type UserList struct {
	Id     int `json:"id"`
	UserId int `json:"user_id"`
	ListId int `json:"listId"`
}
type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
type ListItem struct {
	Id     int `json:"id"`
	ListId int `json:"listId"`
	ItemId int `json:"itemId"`
}
