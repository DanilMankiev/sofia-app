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
		categorys := api.Group("/category")
		{
			categorys.POST("/", h.createCategory)
			categorys.GET("/", h.getAllCategory)
			categorys.GET("/:id", h.getCategoryById)
			categorys.PUT("/:id", h.updateCategory)
			categorys.DELETE("/:id", h.deleteCategory)

			products := categorys.Group(":id/products")
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

		reviews := api.Group("review")
		{
			reviews.POST("/", h.createReview)
			reviews.GET("/", h.getAllReview)
			reviews.PUT("/:id", h.updateReview)
			reviews.DELETE("/:id", h.deleteReview)
		}

		blog := api.Group("blog")
		{
			blog.POST("/", h.createBlog)
			blog.GET("/", h.getAllBlog)
			blog.GET("/:id", h.getBlogById)
			blog.PUT("/:id", h.updateBlog)
			blog.DELETE("/:id", h.deleteBlog)

			images := blog.Group(":id/images")
			{
				images.POST("/", h.createBlogImage)
				images.GET("/", h.getAllBlogImages)
				images.GET("/:im_id", h.getBlogImageById)
				images.PUT("/:im_id", h.updateBlogImage)
				images.DELETE("/:im_id", h.deleteBlogImage)
			}

		}
		api.Static("/stat-images","./images")
		api.Static("/image","./image")
		
	}
	return router
}
