package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/book", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"Title": "Book",
				"Year":  "2024",
				"Genre": "Drama",
			})
		})
		v1.POST("/book", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"Title": "Book",
				"Year":  "2024",
				"Genre": "Drama",
			})
		})
		v1.DELETE("/book", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"Title": "Book",
				"Year":  "2024",
				"Genre": "Drama",
			})
		})
	}
}
