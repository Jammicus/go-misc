package main

import (
	"os"
	"strconv"
)

func main() {

	upto, _ := strconv.ParseInt(os.Args[1], 10, 64)
	fibonacci(upto)
}

func fibonacci(upto int64) int64 {
	number := int64(0)
	numberOne := int64(0)
	numberTwo := int64(0)

	if upto == 0 || upto == 1 {
		number = upto
		upto = 0
	}

	for i := int64(0); i < upto; i++ {
		if i == 0 {
			numberOne = 0
			numberTwo = 0
			number = numberOne + numberTwo

		}
		if i == 1 {
			numberOne = 1
			numberTwo = 0
			number = numberOne + numberTwo

		}
		if i > 1 {
			numberTwo = numberOne
			numberOne = number
			number = numberOne + numberTwo
		}
		println(number)
	}
	return number
}
