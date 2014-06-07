package main

import (
	"fmt"
)

func main() {
	var multiples []int
	for i := 1; i < 1000; i++ {
		if i%5 == 0 || i%3 == 0 {
			multiples = append(multiples, i)
		}
	}
	result := 0
	for _, e := range multiples {
		result += e
	}
	fmt.Println(result)
}
