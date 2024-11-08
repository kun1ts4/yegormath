package main

func Determinant(matrix Matrix) int {
	if matrix.size == 1 {
		return determinant1x1(matrix)
	} else if matrix.size == 2 {
		return determinant2x2(matrix)
	} else if matrix.size == 3 {
		return determinant3x3(matrix)
	} else {
		det := 0
		for j := 0; j < matrix.size; j++ {
			if j%2 == 0 {
				det += matrix.numbers[0][j] * Determinant(Minor(matrix, 0, j))
			} else {
				det += matrix.numbers[0][j] * -1 * Determinant(Minor(matrix, 0, j))
			}
		}
		return det
	}
}

func determinant1x1(matrix Matrix) int {
	return matrix.numbers[0][0]
}

func determinant2x2(matrix Matrix) int {
	return matrix.numbers[0][0]*matrix.numbers[1][1] - matrix.numbers[0][1]*matrix.numbers[1][0]
}

func determinant3x3(matrix Matrix) int {
	general := matrix.numbers[0][0]*matrix.numbers[1][1]*matrix.numbers[2][2] + matrix.numbers[1][0]*matrix.numbers[2][1]*matrix.numbers[0][2] + matrix.numbers[0][1]*matrix.numbers[2][0]*matrix.numbers[1][2]
	secondary := matrix.numbers[0][2]*matrix.numbers[1][1]*matrix.numbers[2][0] + matrix.numbers[0][0]*matrix.numbers[1][2]*matrix.numbers[2][1] + matrix.numbers[0][1]*matrix.numbers[1][0]*matrix.numbers[2][2]
	return general - secondary
}
