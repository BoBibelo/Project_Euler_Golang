/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/

package main

import "fmt"
import "math"

var MAX int = 2000000

func main() {
	ans := 2
	for i := 3; i < MAX; i += 2 {
		if isPrime(float64(i)) {
			ans += i
		}
	}
	fmt.Println("Answer is ", ans)
}

func isPrime(nbr float64) bool {
	for i := math.Trunc(math.Sqrt(nbr)) + 1; i >= 2; i-- {
		var tmp float64 = nbr / i
		if tmp == math.Trunc(tmp) {
			return false
		}
	}
	return true
}
