package handler

import (
	"trenlly/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {

	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}
	api := router.Group("/api", h.UserIdentity)

	{
		tren := api.Group("/tren")
		{
			tren.POST("/", h.CreateTren)
			tren.GET("/", h.GetAllTren)
			tren.GET("/:id", h.GetTrenById)
			tren.PUT("/:id", h.UpdateTren)
			tren.DELETE("/:id", h.DeleteTren)
		}

	}

	return router

}
