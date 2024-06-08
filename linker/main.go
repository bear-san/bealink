package main

import (
	"github.com/bear-san/bealink/linker/internal/link"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/:path", link.Handle)

	err := server.Run("0.0.0.0:8001")
	if err != nil {
		panic(err)
	}
}
