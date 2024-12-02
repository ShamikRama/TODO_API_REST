package handlers

import (
	"net/http"

	"TODO_APP/internal/model"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input model.TodoList

	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	listId, err := h.services.TodoList.Create(userId, input)
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": listId,
	})

}

func (h *Handler) getAllList(c *gin.Context) {

}

func (h *Handler) deleteListByID(c *gin.Context) {

}

func (h *Handler) updateListById(c *gin.Context) {

}

func (h *Handler) getListById(c *gin.Context) {

}
