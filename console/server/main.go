package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()

	err := server.Run("0.0.0.0:8000")
	if err != nil {
		panic(err)
	}
}
