package sieve

import (
	"log"
	"math"
)

type sieveObject struct{}

type Sieve interface {
	NthPrime(n int64) int64
}

func NewSieve() Sieve {
	var sieve sieveObject
	return sieve
}

func (m sieveObject) NthPrime(n int64) int64 {
	if n < 0 {
		log.Fatal("negative index not allowed")
	}
	primeList := make([]int64, n+1)
	primeList[0] = 2
	count := int64(1)

	for num := int64(3); ; {
		if count <= n {
			if !isCandidateDivisible(primeList, num) {
				primeList[count] = num
				count++
			}
			num += 2 // skip even numbers
		} else {
			break
		}
	}
	return primeList[n]
}

func isCandidateDivisible(primeList []int64, num int64) bool {
	for i := 0; i < len(primeList); i++ {
		if math.Sqrt(float64(num)) < float64(primeList[i]) {
			return false
		}
		if num%primeList[i] == 0 {
			return true
		}
	}
	return false
}
