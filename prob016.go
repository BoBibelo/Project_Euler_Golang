/*
2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2^1000?
*/

package main

import (
	"fmt"
)

var POWER int = 1000

func main() {
	pow := []int64{2}
	for i := 1; i < POWER; i++ {
		pow = add(pow, pow)
	}

	sum := sum_digits(pow)
	/*
		fmt.Print("2^", POWER, " = ")
		display_x_digits(pow, len(pow))
		fmt.Println()
	*/
	fmt.Println("Sum of its digits is", sum)
}

func sum_digits(nbr []int64) int64 {
	var sum int64
	for i := 0; i < len(nbr); i++ {
		sum += nbr[i]
	}
	return sum
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

func display_x_digits(nbr []int64, x int) {
	for i := len(nbr) - 1; i >= 0 && i > len(nbr)-1-x; i-- {
		fmt.Print(nbr[i])
	}
}
