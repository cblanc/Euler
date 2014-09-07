package main

import "fmt"

func main() {
	nth := 10001

	primes := []int{2}

	current := 2

	for len(primes) < nth {
		isPrime := true
		for _, prime := range primes {
			// If divisible, select next current and mark current as not prime
			if current%prime == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, current)
		}
		current++
	}

	fmt.Println(primes[len(primes)-1])
}
