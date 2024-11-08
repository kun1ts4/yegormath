package main

import (
	"fmt"
)

func main() {
	// Main function initializes the matrix and prints its determinant.
	matrix := [][]int{
		{3, -3, -5, 8},
		{-3, 2, 4, -6},
		{2, -5, -7, 5},
		{-4, 3, 5, -6},
	}
	fmt.Println(Determinant(matrix))
}
