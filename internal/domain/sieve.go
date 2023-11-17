package domain

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

func NthPrime(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("negative value not allowed")
	}
	if n >= 1e7 {
		gin.DefaultWriter.Write([]byte(fmt.Sprintf("Warning: algorithm may take awhile for inputs larger than 1e7. input: %v\n", n)))
	}
	primeList := make([]int, 0, n)
	primeList = append(primeList, 2)
	num := 3

	for len(primeList) <= n {
		if isPrime(num, primeList) {
			primeList = append(primeList, num)
		}
		num += 2 // skip even numbers
	}
	return primeList[n], nil
}

func isPrime(num int, primeList []int) bool {
	for _, p := range primeList {
		if num < p*p {
			return true
		}
		if num%p == 0 {
			return false
		}
	}
	return true
}
