package matrix

func Minor(matrix Matrix, x int, y int) Matrix {
	mnr := Matrix{Size: matrix.Size - 1}
	for i := 0; i < matrix.Size; i++ {
		var temp []int
		for j := 0; j < matrix.Size; j++ {
			if j != y {
				temp = append(temp, matrix.Numbers[i][j])
			}
		}
		if i != x {
			mnr.Numbers = append(mnr.Numbers, temp)
		}
	}
	return mnr
}
