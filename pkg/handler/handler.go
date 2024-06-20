package handler

import (
	"market/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
	hub      *Hub
}

func NewHandler(services *service.Service, h *Hub) *Handler {
	return &Handler{
		services: services,
		hub:      h,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/", h.GetProducts)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api", h.UserIdentity)
	{
		products := api.Group("/products")
		{
			products.POST("/createProduct", h.CreateProduct)
			products.GET("/buy/:prodId/:count", h.BuyProduct)

		}
	}
	return router
}
