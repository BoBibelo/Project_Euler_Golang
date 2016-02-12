/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/

package main

import "fmt"
import "math"

var NBR float64 = 600851475143

func main() {
	for k := math.Trunc(math.Sqrt(NBR)) + 1; k >= 2; k-- {
		var tmp float64 = NBR / k
		if tmp == math.Trunc(tmp) && isPrime(k) {
			fmt.Println("Largest prime factor = ", k)
			return
		}
	}

	fmt.Println("Error.")
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
