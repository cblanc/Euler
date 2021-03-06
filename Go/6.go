// The sum of the squares of the first ten natural numbers is,

// 12 + 22 + ... + 102 = 385
// The square of the sum of the first ten natural numbers is,

// (1 + 2 + ... + 10)2 = 552 = 3025
// Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

package main

import (
	"fmt"
	"math"
)

func intArray(n int) (result []int) {
	for i := 1; i <= n; i++ {
		result = append(result, i)
	}
	return result
}

func sumOfSquare(numbers []int) (result int) {
	for _, number := range numbers {
		result += int(math.Pow(float64(number), 2))
	}
	return result
}

func squareOfSum(numbers []int) (result int) {
	for _, number := range numbers {
		result += number
	}
	return int(math.Pow(float64(result), 2))
}

func main() {
	set := intArray(100)
	fmt.Println(squareOfSum(set) - sumOfSquare(set))
}
