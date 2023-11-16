package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gjbranham/eratosthenes/internal/config"
	"github.com/gjbranham/eratosthenes/internal/handler"
)

func RunServer() {
	cfg := config.LoadConfig()

	router := gin.Default()

	router.SetTrustedProxies([]string{"localhost"})

	router.GET("/nthPrime/:n", handler.GetPrime)

	router.Run(cfg.Host + ":" + cfg.Port)
}
