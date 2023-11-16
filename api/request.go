package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gjbranham/specter-ops/internal/sieve"
)

type response struct {
	Prime int64 `json:"primeNum"`
}

func GetPrime(c *gin.Context) {
	param := c.Param("n")
	primeIdx, err := strconv.Atoi(param) // this will fail if n > sizeof int so casting to int64 later is fine
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter must be an integer"})
		return
	}
	if primeIdx < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter must be greater than 0"})
	}

	sieve := sieve.NewSieve()
	prime := sieve.NthPrime(int64(primeIdx))

	response := response{
		Prime: prime,
	}
	c.IndentedJSON(http.StatusOK, response)
}
