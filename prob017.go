/*
If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.

If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?

NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23 letters and 115 (one hundred and fifteen) contains 20 letters. The use of "and" when writing out numbers is in compliance with British usage.
*/

package main

import (
	"fmt"
)

func main() {
	uni := []int{0, 3, 3, 5, 4, 4, 3, 5, 5, 4, 3, 6, 6, 8, 8, 7, 7, 9, 8, 8}
	diz := []int{0, 3, 6, 6, 5, 5, 5, 7, 6, 6}
	hun := 7
	tho := 8
	and := 3

	var ans int
	for i := 1; i < 1000; i++ {
		u := i % 10
		d := (i%100 - u) / 10
		h := ((i % 1000) - (d * 10) - u) / 100

		if h != 0 {
			ans += uni[h] + hun
			if d != 0 || u != 0 {
				ans += and
			}
		}
		if d == 0 || d == 1 {
			ans += uni[d*10+u]
		} else {
			ans += diz[d] + uni[u]
		}
	}
	ans += uni[1] + tho

	fmt.Println("Number of letter is", ans)
}
