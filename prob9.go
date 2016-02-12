/*
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
a^2 + b^2 = c^2

For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/

package main

import "fmt"
import "math"

var OBJ float64 = 1000

func main() {
	for a := OBJ; a >= 1; a-- {
		for b := OBJ; b >= 1; b-- {
			for c := OBJ; c >= 1; c-- {
				if a+b+c == OBJ {
					if math.Pow(a, 2)+math.Pow(b, 2) == math.Pow(c, 2) {
						fmt.Println("a*b*c =", int64(a*b*c))
						return
					}
				}
			}
		}
	}
	fmt.Println("oops.")
}
