package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gjbranham/eratosthenes/internal/config"
	"github.com/gjbranham/eratosthenes/internal/handler"
)

func RunServer() {
	cfg := config.LoadConfig() // get host and port from config file; default otherwise

	router := gin.Default()

	router.SetTrustedProxies([]string{"localhost"})

	router.GET("/nthPrime/:n", handler.GetPrime) // define simple GET endpoint

	router.Run(cfg.Host + ":" + cfg.Port) // run server
}
