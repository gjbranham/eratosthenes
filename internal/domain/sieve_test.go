package domain

import (
	"fmt"
	"log"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNthPrimeTable(t *testing.T) {
	type test struct {
		n     int64
		prime int64
	}

	tests := []test{
		{
			n: 0, prime: 2,
		},
		{
			n: 19, prime: 71,
		},
		{
			n: 99, prime: 541,
		},
		{
			n: 500, prime: 3581,
		},
		{
			n: 986, prime: 7793,
		},
		{
			n: 2000, prime: 17393,
		},
		{
			n: 100000, prime: 1299721,
		},
		{
			n: 1000000, prime: 15485867,
		},
		// {
		// 	n: 10000000, prime: 179424691,
		// }, // takes 10-15 seconds
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("n: %v, prime: %v", tt.n, tt.prime), func(t *testing.T) {
			prime, err := NthPrime(tt.n)
			if err != nil {
				t.Fatalf("unexpected error during test: %v", err)
			}
			assert.Equal(t, prime, tt.prime)
		})
	}
}

func TestNegative(t *testing.T) {
	_, err := NthPrime(-1)
	if err == nil {
		t.Fatal("expected error during test")
	}
}

func FuzzNthPrime(f *testing.F) {
	f.Fuzz(func(t *testing.T, n int64) {
		val, err := NthPrime(n)
		if err != nil {
			log.Printf("fuzz test skipping val %v due to err: %v\n", n, err)
			return
		}
		if !big.NewInt(val).ProbablyPrime(0) {
			t.Errorf("the sieve produced a non-prime number at index %d", n)
		}
	})
}
