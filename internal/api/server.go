package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gjbranham/eratosthenes/internal/config"
	"github.com/gjbranham/eratosthenes/internal/handler"
)

func RunServer() {
	cfg := config.LoadConfig() // load server and host from config file if available

	router := gin.Default()

	router.SetTrustedProxies([]string{"localhost"})

	router.GET("/nthPrime/:n", handler.GetPrime) // define simple GET endpoint

	router.Run(cfg.Host + ":" + cfg.Port) // run server
}
