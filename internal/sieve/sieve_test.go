package sieve

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNthPrime(t *testing.T) {
	sieve := NewSieve()

	assert.Equal(t, int64(2), sieve.NthPrime(0))
	assert.Equal(t, int64(71), sieve.NthPrime(19))
	assert.Equal(t, int64(541), sieve.NthPrime(99))
	assert.Equal(t, int64(3581), sieve.NthPrime(500))
	assert.Equal(t, int64(7793), sieve.NthPrime(986))
	assert.Equal(t, int64(17393), sieve.NthPrime(2000))
	assert.Equal(t, int64(1299721), sieve.NthPrime(100000))
	assert.Equal(t, int64(15485867), sieve.NthPrime(1000000))
	// assert.Equal(t, 179424691, sieve.NthPrime(10000000)) // about 9 seconds
}

func FuzzNthPrime(f *testing.F) {
	sieve := NewSieve()

	f.Fuzz(func(t *testing.T, n int64) {
		if n < 0 {
			t.Logf("received a negative input: %v\nsieving method exits on negative inputs", n)
			return
		}
		if !big.NewInt(sieve.NthPrime(n)).ProbablyPrime(0) {
			t.Errorf("the sieve produced a non-prime number at index %d", n)
		}
	})
}