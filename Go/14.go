// The following iterative sequence is defined for the set of positive integers:

// n → n/2 (n is even)
// n → 3n + 1 (n is odd)

// Using the rule above and starting with 13, we generate the following sequence:

// 13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
// It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

// Which starting number, under one million, produces the longest chain?

// NOTE: Once the chain starts the terms are allowed to go above one million.

package main

import "fmt"

func even(n int) int {
	return n / 2
}

func odd(n int) int {
	return (3 * n) + 1
}

func next(n int) int {
	if n%2 == 0 {
		return even(n)
	} else {
		return odd(n)
	}
}

type collatz struct {
	startingNumber int
}

func (c collatz) chainLength() int {
	chainLength := 1
	start := c.startingNumber
	for start != 1 {
		start = next(start)
		chainLength++
	}
	return chainLength
}

func main() {
	start := 1
	maxCollatz := 1000000
	longestChain := 0
	var longestCollatz collatz
	var c collatz

	for i := start; i < maxCollatz; i++ {
		c = collatz{
			startingNumber: i,
		}
		if length := c.chainLength(); length > longestChain {
			longestCollatz = c
			longestChain = length
		}
	}
	fmt.Println("Starting Number of Longest Chain:")
	fmt.Println(longestCollatz.startingNumber)
}
