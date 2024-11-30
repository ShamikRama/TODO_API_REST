package handler

import (
	"TODO_APP/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type signInRequest struct {
	username string `json:"username" binding:"required"`
	password string `json:"password" binding:"required"`
}

func (h *Handler) SignUp(c *gin.Context) {
	var input model.User
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	id, err := h.services.Authorization.Create(input)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) SignIn(c *gin.Context) {
	var input signInRequest

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// TODO : generate token
	// TODO : parsing token
}
