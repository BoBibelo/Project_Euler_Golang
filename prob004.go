/*
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/

package main

import "fmt"
import "strconv"
import "math"

func main() {
	var ans float64 = -1

	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			if isPalindrome(strconv.Itoa(i * j)) {
				ans = math.Max(ans, float64(i*j))
			}
		}
	}

	fmt.Println("Answer is ", ans)
}

func isPalindrome(str string) bool {
	if str == reverseString(str) {
		return true
	} else {
		return false
	}
}

func reverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
