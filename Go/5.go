/*

2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

*/

package main

import (
	"fmt"
)

const (
	interval = 2 * 2 * 2 * 2 * 3 * 3 * 5 * 7 * 11 * 13 * 17 * 19
)

func lcm(n int) int {
	factors := []int{}

	return len(factors)
}

func main() {
	// i := 1000000
	// fmt.Println(i)
	// for {
	// 	fmt.Println(i)
	// 	if isDivisible(i) {
	// 		fmt.Println(i)
	// 	}
	// 	i += interval
	// }
	fmt.Println(interval)
}

// For numbers 1 to 20
// Break down into factors
// Are all factors in map
