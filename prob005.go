/*
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

package main

import "fmt"

func main() {
	list := []int{3, 6, 7, 8, 9, 11, 13, 14, 15, 16, 17, 18, 19, 20}
	var i int = 20
	for true {
		if isDivisible(i, list) {
			fmt.Println("Answer is ", i)
			return
		} else {
			i += 20
		}
	}
}

func isDivisible(nbr int, list []int) bool {
	for _, div := range list {
		if nbr%div != 0 {
			return false
		}
	}
	return true
}
