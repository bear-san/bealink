package main

import (
	"github.com/bear-san/bealink/console/server/internal/auth"
	"github.com/bear-san/bealink/console/server/internal/link"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	server := gin.Default()

	apiGroup := server.Group("/api")
	linkGroup := apiGroup.Group("/links")
	linkGroup.GET("/", link.List)
	linkGroup.POST("/", link.Create)
	linkGroup.DELETE("/:lid", link.Delete)

	authGroup := apiGroup.Group("/auth")
	authGroup.GET("/login", auth.Login)
	authGroup.GET("/callback", auth.Callback)

	apiGroup.GET("/metadata", func(req *gin.Context) {
		req.JSON(200, gin.H{
			"link_host": os.Getenv("LINK_HOST"),
		})
	})

	err := server.Run("0.0.0.0:8000")
	if err != nil {
		panic(err)
	}
}
