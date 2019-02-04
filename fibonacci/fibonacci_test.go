package main

import "testing"

func TestFibonacci(t *testing.T) {
	var testcases = []struct {
		input    int64
		expected int64
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{14, 377},
		{50, 12586269025},
		{90, 2880067194370816120},
	}

	for _, test := range testcases {
		if item := fibonacci(test.input); item != test.expected {
			t.Errorf("fibonacci(%d) = %v", test.input, item)
			print("Here", test.input, item)
		}
	}
}
