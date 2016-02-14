/*
Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.

For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.
*/

package main

import (
	"fmt"
)

var CST int = 10000

func main() {
	var sum int
	var excluded []int

	for i := 1; i < CST; i++ {
		if is_in(excluded, i) {
			continue
		}
		var tmp = d(i)
		if tmp != i && d(tmp) == i {
			sum += i + tmp
			excluded = append(excluded, tmp, i)
		}
	}

	fmt.Println("Answer is", sum)
}

func is_in(arr []int, i int) bool {
	for _, x := range arr {
		if x == i {
			return true
		}
	}
	return false
}

func d(nbr int) int {
	var sum int

	for i := 1; i < nbr; i++ {
		if nbr%i == 0 {
			sum += i
		}
	}

	return sum
}
