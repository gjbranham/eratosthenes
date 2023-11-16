package domain

import (
	"errors"
)

func NthPrime(n int64) (int64, error) {
	if n < 0 {
		return 0, errors.New("negative value not allowed")
	}
	primeList := make([]int64, 0, n+1)
	primeList = append(primeList, 2)
	count, num := int64(1), int64(3)

	for count <= n {
		if isPrime(num, primeList) {
			primeList = append(primeList, num)
			count++
		}
		num += 2 // skip even numbers
	}
	return primeList[n], nil
}

func isPrime(num int64, primeList []int64) bool {
	for _, p := range primeList {
		if p*p > num {
			return true
		}
		if num%p == 0 {
			return false
		}
	}
	return true
}
