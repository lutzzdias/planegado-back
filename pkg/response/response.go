package response

import "github.com/gin-gonic/gin"

func Success(context *gin.Context, data any) {
	context.JSON(200, gin.H{
		"status": "success",
		"data":   data,
	})
}

func Error(context *gin.Context, status int, message string) {
	context.JSON(status, gin.H{
		"status":  "error",
		"message": message,
	})
}
