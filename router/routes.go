package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rphmauriciodev/crud-GO.git/handler"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/book", handler.GetBookHandler)
		v1.POST("/book", handler.CreateBookHandler)
		v1.PUT("/book", handler.UpdateBookHandler)
		v1.DELETE("/book", handler.DeleteBookHandler)
	}
}
