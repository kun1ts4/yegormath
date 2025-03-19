package matrix

func Determinant(matrix Matrix) int {
	if matrix.Size == 1 {
		return determinant1x1(matrix)
	} else if matrix.Size == 2 {
		return determinant2x2(matrix)
	} else if matrix.Size == 3 {
		return determinant3x3(matrix)
	} else {
		det := 0
		for j := 0; j < matrix.Size; j++ {
			if j%2 == 0 {
				det += matrix.Numbers[0][j] * Determinant(Minor(matrix, 0, j))
			} else {
				det += -1 * matrix.Numbers[0][j] * Determinant(Minor(matrix, 0, j))
			}
		}
		return det
	}
}

func determinant1x1(matrix Matrix) int {
	return matrix.Numbers[0][0]
}

func determinant2x2(matrix Matrix) int {
	return matrix.Numbers[0][0]*matrix.Numbers[1][1] - matrix.Numbers[0][1]*matrix.Numbers[1][0]
}

func determinant3x3(matrix Matrix) int {
	general := matrix.Numbers[0][0]*matrix.Numbers[1][1]*matrix.Numbers[2][2] +
		matrix.Numbers[1][0]*matrix.Numbers[2][1]*matrix.Numbers[0][2] +
		matrix.Numbers[0][1]*matrix.Numbers[2][0]*matrix.Numbers[1][2]
	secondary := matrix.Numbers[0][2]*matrix.Numbers[1][1]*matrix.Numbers[2][0] +
		matrix.Numbers[0][0]*matrix.Numbers[1][2]*matrix.Numbers[2][1] +
		matrix.Numbers[0][1]*matrix.Numbers[1][0]*matrix.Numbers[2][2]
	return general - secondary
}
