package handler

import (
	"net/http"
	reapeatTodo "repeat/Todo"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createList(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}
	var input reapeatTodo.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err = h.services.TodoList.Create(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllLists(c *gin.Context) {

}
func (h *Handler) getListById(c *gin.Context) {

}
func (h *Handler) updateList(c *gin.Context) {

}
func (h *Handler) deleteList(c *gin.Context) {

}
