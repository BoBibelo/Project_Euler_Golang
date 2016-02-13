/*
You are given the following information, but you may prefer to do some research for yourself.

1 Jan 1900 was a Monday.
Thirty days has September,
April, June and November.
All the rest have thirty-one,
Saving February alone,
Which has twenty-eight, rain or shine.
And on leap years, twenty-nine.
A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.

How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
*/

package main

import (
	"fmt"
)

func main() {
	var sundays int
	var cursor int = 1
	months_norm := []int{1, 32, 60, 91, 121, 152, 182, 213, 244, 274, 305, 335}
	months_leap := []int{1, 32, 61, 92, 122, 153, 183, 214, 245, 275, 306, 336}

	for year := 1900; year < 2001; year++ {
		var days int
		var months []int

		if year%4 == 0 && year != 1900 {
			days = 366
			months = months_leap
		} else {
			days = 365
			months = months_norm
		}

		for i := 1; i <= days; i, cursor = i+1, cursor+1 {
			if cursor == 7 {
				cursor = 0
				if year != 1900 && is_in(months, i) {
					sundays++
				}
			}
		}
	}

	fmt.Println("There are", sundays, "sundays")
}

func is_in(months []int, day int) bool {
	for _, i := range months {
		if i == day {
			return true
		}
	}
	return false
}
