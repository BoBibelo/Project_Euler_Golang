/*
Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down, there are exactly 6 routes to the bottom right corner.
How many such routes are there through a 20×20 grid?
*/

package main

import (
	"fmt"
)

var size_grid int = 20

func main() {
	grid := make([][]int, size_grid+1)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, size_grid+1)
	}

	for i := 0; i <= size_grid; i++ {
		for j := 0; j <= size_grid; j++ {
			if j == 0 || i == 0 {
				grid[i][j] = 1
			} else {
				grid[i][j] = grid[i-1][j] + grid[i][j-1]
			}
		}
	}

	fmt.Println("Number of path is", grid[size_grid][size_grid])
}
