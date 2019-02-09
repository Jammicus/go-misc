package main

import (
	"math/big"
	"testing"
)

func TestIsPrime(t *testing.T) {
	var testcases = []struct {
		number int64
		prime  bool
	}{
		{0, false},
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{8, false},
		{9, false},
		{100, false},
		{101, true},
		{109, true},
		{110, false},
		{150, false},
		{151, true},
		{163, true},
	}
	for _, test := range testcases {
		num := big.NewInt(int64(test.number))
		if item := isPrime(*num); item != test.prime {
			t.Errorf("isPrime(%v) = %v", test.number, item)
		}
	}
}

func BenchmarkIsItPrime(b *testing.B) {

	for n := 0; n < b.N; n++ {
		i := big.NewInt(int64(n))
		isPrime(*i)
	}
}
