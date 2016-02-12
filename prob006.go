/*
The sum of the squares of the first ten natural numbers is,
1^2 + 2^2 + ... + 10^2 = 385

The square of the sum of the first ten natural numbers is,
(1 + 2 + ... + 10)^2 = 55^2 = 3025

Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/

package main

import "fmt"
import "math"

var NBR float64 = 100

/*
**	Simple math formulas
**	(sigma i from 1 to NBR) power 2 - (sigma i power 2 from 1 to NBR)
 */
func main() {
	var ans float64
	ans += math.Pow(NBR, 4) / 4
	ans += math.Pow(NBR, 3) / 6
	ans -= math.Pow(NBR, 2) / 4
	ans -= NBR / 6

	fmt.Println("Answer is", int64(ans))
}
