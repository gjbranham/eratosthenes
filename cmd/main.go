package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gjbranham/specter-ops/api"
)

func main() {
	// Set up the server here since it's not a lot of code
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.GET("/prime/:n", api.GetPrime)

	router.Run("localhost:3000")
}
