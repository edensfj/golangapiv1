package api

import (
	"back_go/handler"

	"github.com/gin-gonic/gin"
)

func Routes(e *gin.Engine) {
	group := e.Group("/api")
	// group.Use(middleware.AuthMiddleware())
	group.GET("/", handler.Hello)
}
