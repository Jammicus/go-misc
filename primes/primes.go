package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {

	numbers := os.Args[1:]

	for i := 0; i < len(numbers); i++ {
		num, err := strconv.Atoi(numbers[i])
		if err != nil {
			panic("Error with input, please ensure its an integer")
		}
		convnum := big.NewInt(int64(num))
		isPrime(*convnum)
		fmt.Println("is", convnum, "a prime number?", isPrime(*convnum))
	}
}

func isPrime(number big.Int) bool {
	return number.ProbablyPrime(100000)
}
