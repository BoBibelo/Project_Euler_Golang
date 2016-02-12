/*
Work out the first ten digits of the sum of the following one-hundred 50-digit numbers.
See prob13.txt
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("prob013.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := []int64{0}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum = add(sum, create_nbr(scanner.Text()))
	}

	fmt.Println("Full number:")
	display_x_digits(sum, len(sum))
	fmt.Println("\nFirst ten digits:")
	display_x_digits(sum, 10)
	fmt.Println()
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

func create_nbr(str string) []int64 {
	nbr := make([]int64, len(str))
	for i, j := len(nbr)-1, 0; i >= 0 && j < len(nbr); i, j = i-1, j+1 {
		nbr[j] = int64(str[i] - '0')
	}
	return nbr
}

func display_x_digits(nbr []int64, x int) {
	for i := len(nbr) - 1; i >= 0 && i > len(nbr)-1-x; i-- {
		fmt.Print(nbr[i])
	}
}
