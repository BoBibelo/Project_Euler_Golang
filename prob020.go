/*
n! means n × (n − 1) × ... × 3 × 2 × 1

For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!
*/

package main

import (
	"fmt"
)

var FACT int = 100

func main() {
	prod := []int64{2}
	for i := 3; i <= FACT; i++ {
		prod = multi(prod, i)
	}
	var sum int64 = sum_digits(prod)

	fmt.Println("The sum of the digits is", sum)
}

func multi(nbr []int64, x int) []int64 {
	cpy := []int64{0}
	for i := 0; i < x; i++ {
		cpy = add(cpy, nbr)
	}
	return cpy
}

func add(old []int64, new []int64) []int64 {
	var ans []int64
	var carry int64

	for i, j := 0, 0; i < len(old) || j < len(new); i, j = i+1, j+1 {
		if i < len(old) {
			carry += old[i]
		}
		if j < len(new) {
			carry += new[j]
		}
		ans = append(ans, carry%10)
		carry /= 10
	}
	if carry != 0 {
		ans = append(ans, carry)
	}
	return ans
}

func sum_digits(nbr []int64) int64 {
	var sum int64
	for _, i := range nbr {
		sum += i
	}
	return sum
}
