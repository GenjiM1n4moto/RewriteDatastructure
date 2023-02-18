package rewritedatastructure

func test1() int {
	i := 7 / 2
	return i
}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	l := len(matrix)
	h := len(matrix[0])
	r := l / 2
	c := h / 2
	var checkMatrix [l][h]bool
	// use make
	for i := 0; i < l; i++ {
		for j := 0; j < h; j++ {
			checkMatrix[i][j] = false
		}
	}
	if matrix[r][c] < target {

	}
}
