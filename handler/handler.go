package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBookHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Title": "Book",
		"Year":  "2024",
		"Genre": "Drama",
	})
}
func GetBookHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Title": "Book",
		"Year":  "2024",
		"Genre": "Drama",
	})
}
func UpdateBookHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Title": "Book",
		"Year":  "2024",
		"Genre": "Drama",
	})
}
func DeleteBookHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Title": "Book",
		"Year":  "2024",
		"Genre": "Drama",
	})
}