/*

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?

*/

package main

import (
	"fmt"
)

// Checks if new prime discovered
func isNewPrime(candidate int, knownPrimes []int) bool {
	for _, e := range knownPrimes {
		if candidate%e == 0 {
			return false
		}
	}
	return true
}

func main() {
	var knownPrimes []int

	n := 600851475143
	limit := n

	for i := 2; i <= limit; i++ {
		if n%i != 0 {
			continue
		}
		if !isNewPrime(i, knownPrimes) {
			continue
		}
		fmt.Println(i)
		knownPrimes = append(knownPrimes, i)
		limit = n / i
	}
	fmt.Println(knownPrimes[len(knownPrimes)-1])
}

// Steps
// Get first number
// Is it a factor? No -> exit
// Is the factor prime? No -> exit
// Add prime to list + update limit
