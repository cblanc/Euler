// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

// Find the sum of all the primes below two million.

package main

import "fmt"

func eliminateFactors(candidates [2000000]bool, factor int) [2000000]bool {
	index := factor * 2
	for index < len(candidates) {
		candidates[index] = false
		index := factor
	}
	return candidates
}

func initCandidates() [2000000]bool {
	var candidates [2000000]bool
	for i, _ := range candidates {
		candidates[i] = true
	}
	return candidates
}

func main() {
	candidates := initCandidates()
	fmt.Println(candidates[0])

	// Initialise the first few
	candidates[0] = false // 0 isn't prime
	candidates[1] = false // 1 isn't prime
	candidates[2] = true
	lastPrime := 2

}

// Naieve Implementation

// func isPrime(n int, primes []int) bool {
// 	for _, prime := range primes {
// 		if prime > n || n%prime == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func nextPrime(primes []int) int {
// 	n := primes[len(primes)-1] + 1
// 	for !isPrime(n, primes) {
// 		n++
// 	}
// 	return n
// }

// func sum(numbers []int) int {
// 	var result int
// 	for _, n := range numbers {
// 		result += n
// 	}
// 	return result
// }

// func main() {
// 	primes := []int{2}
// 	max := 2000000
// 	var next int
// 	for primes[len(primes)-1] < max {
// 		next = nextPrime(primes)
// 		if next < max {
// 			primes = append(primes, nextPrime(primes))
// 		} else {
// 			break
// 		}
// 	}
// 	fmt.Println(sum(primes))
// }