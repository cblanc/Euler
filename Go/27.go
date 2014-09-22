// Euler discovered the remarkable quadratic formula:

// n² + n + 41

// It turns out that the formula will produce 40 primes for the consecutive values n = 0 to 39. However, when n = 40, 402 + 40 + 41 = 40(40 + 1) + 41 is divisible by 41, and certainly when n = 41, 41² + 41 + 41 is clearly divisible by 41.

// The incredible formula  n² − 79n + 1601 was discovered, which produces 80 primes for the consecutive values n = 0 to 79. The product of the coefficients, −79 and 1601, is −126479.

// Considering quadratics of the form:

// n² + an + b, where |a| < 1000 and |b| < 1000

// where |n| is the modulus/absolute value of n
// e.g. |11| = 11 and |−4| = 4
// Find the product of the coefficients, a and b, for the quadratic expression that produces the maximum number of primes for consecutive values of n, starting with n = 0.

package main

import (
	"fmt"
)

type QuadraticCombination struct {
	a                 int
	b                 int
	consecutivePrimes int
}

func isPrime(n int, primes []int) bool {
	limit := n

	for i := 0; primes[i] < limit; i++ {
		if n%primes[i] == 0 {
			return false
		}
		limit = (n / primes[i]) + 1
	}
	return true
}

func maxConsecutivePrimes(a int, b int, primes []int) int {
	result := 0
	for n := 0; ; n++ {
		if isPrime((n*n)+(a*n)+b, primes) {
			result++
		} else {
			break
		}
	}
	return result
}

func main() {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293, 307, 311, 313, 317, 331, 337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409, 419, 421, 431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509, 521, 523, 541, 547, 557, 563, 569, 571, 577, 587, 593, 599, 601, 607, 613, 617, 619, 631, 641, 643, 647, 653, 659, 661, 673, 677, 683, 691, 701, 709, 719, 727, 733, 739, 743, 751, 757, 761, 769, 773, 787, 797, 809, 811, 821, 823, 827, 829, 839, 853, 857, 859, 863, 877, 881, 883, 887, 907, 911, 919, 929, 937, 941, 947, 953, 967, 971, 977, 983, 991, 997, 1009, 1013, 1019, 1021, 1031, 1033, 1039, 1049, 1051, 1061, 1063, 1069, 1087, 1091, 1093, 1097, 1103, 1109, 1117, 1123, 1129, 1151, 1153, 1163, 1171, 1181, 1187, 1193, 1201, 1213, 1217, 1223, 1229, 1231, 1237, 1249, 1259, 1277, 1279, 1283, 1289, 1291, 1297, 1301, 1303, 1307, 1319, 1321, 1327, 1361, 1367, 1373, 1381, 1399, 1409, 1423, 1427, 1429, 1433, 1439, 1447, 1451, 1453, 1459, 1471, 1481, 1483, 1487, 1489, 1493, 1499, 1511, 1523, 1531, 1543, 1549, 1553, 1559, 1567, 1571, 1579, 1583}

	max := QuadraticCombination{
		a:                 0,
		b:                 0,
		consecutivePrimes: 0,
	}

	var maxPrimes int

	for a := -1000; a <= 1000; a++ {
		for b := -1000; b <= 1000; b++ {
			if maxPrimes = maxConsecutivePrimes(a, b, primes); maxPrimes > max.consecutivePrimes {
				max.a = a
				max.b = b
				max.consecutivePrimes = maxPrimes
			}
		}
	}

	fmt.Printf("a:%d\n", max.a)
	fmt.Printf("b:%d\n", max.b)
	fmt.Printf("Max Consecutives:%d\n", max.consecutivePrimes)
	fmt.Printf("Product:%d\n", max.a*max.b)

}
