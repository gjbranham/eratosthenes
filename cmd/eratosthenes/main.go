package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/gjbranham/specter-ops/api"
)

var (
	host = flag.String("host", "localhost", "The server host")
	port = flag.String("port", "3000", "The server port")
)

func main() {
	flag.Parse()
	// Set up the server here since it's not a lot of code
	router := gin.Default()
	router.SetTrustedProxies([]string{"localhost"})

	router.GET("/nthPrime/:n", api.GetPrime)

	router.Run(*host + ":" + *port)
}
