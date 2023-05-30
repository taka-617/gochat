package route

import (
	"github.com/gin-gonic/gin"
	"github.com/taka-617/gochat/controller"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("view/*html")
	r.GET("/", controller.ShowTop)
	r.GET("/auth/login/:provider", controller.AuthLogin)
	r.GET("/auth/callback/:provider", controller.AuthCallback)
	return r
}