package handler

import (
	"github.com/DanilMankiev/sofia-app/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandeler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/list")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllList)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			products := lists.Group(":id/products")
			{
				products.GET("/", h.getAllProducts)
				products.POST("/", h.createProduct)
			}
		}
		products := api.Group("product")
		{
			products.GET("/:id", h.getItemById)
			products.PUT("/:id", h.updateItem)
			products.DELETE("/:id", h.deleteItem)

			images := products.Group(":id/images")
			{
				images.POST("/", h.createImage)
				images.GET("/", h.getAllImages)
				images.GET("/:im_id", h.getImageById)
				images.PUT("/:im_id", h.updateImage)
				images.DELETE("/:im_id", h.deleteImage)
			}
		}

	}
	return router
}
