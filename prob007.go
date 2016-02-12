/*
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?
*/

package main

import "fmt"
import "math"

var NBR int = 10001

func main() {
	var count int = 1
	for i := 2; true; i++ {
		if isPrime(float64(i)) {
			count++
			if count == NBR {
				fmt.Println("Answer is", i)
				return
			}
		}

	}
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
