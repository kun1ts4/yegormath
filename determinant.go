package main

func Determinant(matrix [][]int) int {
	size := len(matrix)
	if size == 1 {
		return determinant1x1(matrix)
	} else if size == 2 {
		return determinant2x2(matrix)
	} else if size == 3 {
		return determinant3x3(matrix)
	} else {
		det := 0
		for j := 0; j < size; j++ {
			if j%2 == 0 {
				det += matrix[0][j] * Determinant(Minor(matrix, 0, j))
			} else {
				det += matrix[0][j] * -1 * Determinant(Minor(matrix, 0, j))
			}
		}
		return det
	}
}

func determinant1x1(matrix [][]int) int {
	return matrix[0][0]
}

func determinant2x2(matrix [][]int) int {
	return matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0]
}

func determinant3x3(matrix [][]int) int {
	general := matrix[0][0]*matrix[1][1]*matrix[2][2] + matrix[1][0]*matrix[2][1]*matrix[0][2] + matrix[0][1]*matrix[2][0]*matrix[1][2]
	secondary := matrix[0][2]*matrix[1][1]*matrix[2][0] + matrix[0][0]*matrix[1][2]*matrix[2][1] + matrix[0][1]*matrix[1][0]*matrix[2][2]
	return general - secondary
}
