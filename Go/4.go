package main

import (
	"fmt"
	"strconv"
)

func isPalendrome(n int) bool {
	number := strconv.Itoa(n)

	length := len(number)

	for i := 0; i <= length / 2; i++ {
		if number[i:i+1] != number[length - i -1:length-i] {
			return false
		}
	}
	return true
}

func main() {
	largestPalendrome := 0

	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			if r := i*j; isPalendrome(r) && (r > largestPalendrome) {
				largestPalendrome = r
			}
		}
	}
	fmt.Println(largestPalendrome)
}