package main

import (
	"fmt"
	"yegormath/internal/matrix"
)

func main() {
	// Main function initializes the matrix and prints its determinant.
	m := matrix.Matrix{
		Numbers: [][]int{
			{5, 9, -10, 0},
			{66, 0, 2, 6},
			{-70, 29, 1, 1},
			{0, 4, -8, 1},
		},
		Size: 4,
	}
	fmt.Println(matrix.Determinant(m))
}
