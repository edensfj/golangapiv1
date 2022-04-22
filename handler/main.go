package handler

import "github.com/gin-gonic/gin"

func Hello(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hola mundo",
	})
}
