package YandexAlgorithms3

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var x, y int
	fmt.Scan(&y, &x)
	coins := make([][]int, y)
	for i := 0; i < y; i++ {
		coins[i] = make([]int, x)
		for j := 0; j < x; j++ {
			fmt.Scan(&coins[i][j])
		}
	}
	maxCoins := make([][]int, y)
	for i := 0; i < y; i++ {
		maxCoins[i] = make([]int, x)
	}
	maxCoins[0][0] = coins[0][0]
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				maxCoins[i][j] = maxCoins[i][j-1] + coins[i][j]
			} else if j == 0 {
				maxCoins[i][j] = maxCoins[i-1][j] + coins[i][j]
			} else {
				maxCoins[i][j] = max(maxCoins[i-1][j], maxCoins[i][j-1]) + coins[i][j]
			}
		}
	}
	fmt.Println(maxCoins[y-1][x-1])
	var path []string
	i := y - 1
	j := x - 1
	for i > 0 || j > 0 {
		if i == 0 {
			path = append(path, "R")
			j--
		} else if j == 0 {
			path = append(path, "D")
			i--
		} else {
			if maxCoins[i-1][j] > maxCoins[i][j-1] {
				path = append(path, "D")
				i--
			} else {
				path = append(path, "R")
				j--
			}
		}
	}
	for i := len(path) - 1; i >= 0; i-- {
		fmt.Print(path[i], " ")
	}
}
