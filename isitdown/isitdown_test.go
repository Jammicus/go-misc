package main

import (
	"testing"
)

func TestIsPrime(t *testing.T) {
	var testcases = []struct {
		website string
		status  bool
	}{
		{"http://fakenews.co.co", false},
		{"https://www.bbc.co.uk/news", true},
		{"https://bbc.co.uk", true},
		{"http://google.co.uk", true},
	}
	for _, test := range testcases {
		if item := isitdown(test.website); item != test.status {
			t.Errorf("isPrime(%v) = %v", test.website, item)
		}

	}
}
