package main

func Minor(matrix [][]int, x int, y int) [][]int {
	mnr := [][]int{}
	for i := 0; i < len(matrix); i++ {
		temp := []int{}
		for j := 0; j < len(matrix); j++ {
			if j != y {
				temp = append(temp, matrix[i][j])
			}
		}
		if i != x {
			mnr = append(mnr, temp)
		}
	}
	return mnr
}
