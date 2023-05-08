package YandexAlgorithms3

import "fmt"

func myMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	doubleD := make([][]int, n)
	for i := 0; i < n; i++ {
		doubleD[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&doubleD[i][j])
		}
	}
	for i := 1; i < n; i++ {
		doubleD[i][0] += doubleD[i-1][0]
	}
	for j := 1; j < m; j++ {
		doubleD[0][j] += doubleD[0][j-1]
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			doubleD[i][j] += myMin(doubleD[i-1][j], doubleD[i][j-1])
		}
	}
	fmt.Println(doubleD[n-1][m-1])
}
