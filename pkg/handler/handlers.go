package handler

import (
	_ "github.com/DanilMankiev/sofia-app/docs"
	"github.com/DanilMankiev/sofia-app/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type Handler struct {
	services *service.Service
}

func NewHandeler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/refresh", h.refresh)
	}
	api := router.Group("/api",h.userIdentity)
	{
		user := api.Group("/user")
		{
			user.GET("/", h.getUser)
			favorites := user.Group("/favorites")
			{
				favorites.POST("/:id", h.createFavorites)
				favorites.GET("/", h.getAllFavorites)
				favorites.DELETE("/:id", h.deleteFavorites)
			}
		}
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
		products := api.Group("products")
		{
			products.GET("/:id", h.getItemById)
			products.PUT("/:id", h.updateItem)
			products.DELETE("/:id", h.deleteItem)

			images := products.Group(":id/images")
			{
				images.POST("/", h.createImage)
				images.DELETE("/", h.deleteImage)
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
				images.DELETE("/", h.deleteBlogImage)
			}

		}
		api.Static("/image", "./image")
	}
	return router
}
