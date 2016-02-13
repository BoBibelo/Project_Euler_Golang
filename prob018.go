/*
By starting at the top of the triangle below and moving to adjacent numbers on the row below, the maximum total from top to bottom is 23.

3
7 4
2 4 6
8 5 9 3

That is, 3 + 7 + 4 + 9 = 23.

Find the maximum total from top to bottom of the triangle below:
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	triangle := create_triangle()

	for len(triangle) > 1 {
		lenght := len(triangle) - 2
		for i := range triangle[lenght] {
			triangle[lenght][i] += max(triangle[lenght+1][i], triangle[lenght+1][i+1])
		}
		triangle = triangle[0 : lenght+1]
	}

	fmt.Println("Answer is", triangle[0][0])
}

func max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func create_triangle() [][]int {
	file, err := os.Open("prob018.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var triangle [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		triangle = append(triangle, convert(scanner.Text()))
	}

	return triangle
}

func convert(line string) []int {
	var line_int []int
	var nbr int

	for i := 0; i < len(line); i = i + 3 {
		nbr = int(line[i]-'0')*10 + int(line[i+1]-'0')
		line_int = append(line_int, nbr)
	}
	return line_int
}
