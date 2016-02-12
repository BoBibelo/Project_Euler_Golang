/*
The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:
13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1

It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.
*/

package main

import (
	"fmt"
)

var CST int = 1000000

func main() {
	var len_chain int = -1
	var nbr int

	for i := 2; i < CST; i++ {
		tmp := prod_chain(i)
		if tmp >= len_chain {
			len_chain = tmp
			nbr = i
		}
	}

	fmt.Println("Longest chain is", nbr, "with a chain lenght of", len_chain)
}

func prod_chain(nbr int) int {
	var lenght int = 1

	for nbr != 1 {
		if nbr%2 == 0 {
			nbr /= 2
		} else {
			nbr = 3*nbr + 1
		}
		lenght++
	}

	return lenght
}
