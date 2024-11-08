package main

func Minor(matrix Matrix, x int, y int) Matrix {
	mnr := Matrix{size: matrix.size - 1}
	for i := 0; i < matrix.size; i++ {
		temp := []int{}
		for j := 0; j < matrix.size; j++ {
			if j != y {
				temp = append(temp, matrix.numbers[i][j])
			}
		}
		if i != x {
			mnr.numbers = append(mnr.numbers, temp)
		}
	}
	return mnr
}
