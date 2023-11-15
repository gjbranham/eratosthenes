package sieve

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNthPrime(t *testing.T) {
	sieve := NewSieve()

	assert.Equal(t, 2, sieve.NthPrime(0))
	assert.Equal(t, 71, sieve.NthPrime(19))
	assert.Equal(t, 541, sieve.NthPrime(99))
	assert.Equal(t, 3581, sieve.NthPrime(500))
	assert.Equal(t, 7793, sieve.NthPrime(986))
	assert.Equal(t, 17393, sieve.NthPrime(2000))
	assert.Equal(t, 15485867, sieve.NthPrime(1000000))
}

func FuzzNthPrime(f *testing.F) {
	sieve := NewSieve()

	f.Fuzz(func(t *testing.T, n int64) {
		if !big.NewInt(sieve.NthPrime(n)).ProbablyPrime(0) {
			t.Errorf("the sieve produced a non-prime number at index %d", n)
		}
	})
}
