package handler

import (
	"TODO_APP/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandle(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up")
		auth.POST("/sign-in")
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/")
			lists.GET("/")
			lists.GET("/:id")
			lists.PUT("/:id")
			lists.DELETE("/:id")

			items := lists.Group(":id/items")
			{
				items.POST("/")
				items.GET("/")
			}
		}

		items := api.Group("items")
		{
			items.GET("/:id")
			items.PUT("/:id")
			items.DELETE("/:id")
		}
	}

	return router
}
