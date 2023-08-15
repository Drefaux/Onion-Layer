package router

import (
	"OnionPractice/app/controller"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	todo := router.Group("/todos")
	{
		todo.GET("/:id", controller.Get)
		todo.POST("", controller.Create)
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	rootPathV2 := "/api-" + os.Getenv("SERVICE_URL") + "/v2"
	category := router.Group(rootPathV2 + "/categories")
	{
		category.GET("/category", controller.GetAllCategories)
		category.POST("", controller.CreateCategory)
	}

	return router
}
