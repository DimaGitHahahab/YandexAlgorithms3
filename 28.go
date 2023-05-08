package YandexAlgorithms3

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	dp[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if n-i > 2 && m-j > 1 {
				dp[i+2][j+1] += dp[i][j]
			}
			if n-i > 1 && m-j > 2 {
				dp[i+1][j+2] += dp[i][j]
			}
		}
	}
	fmt.Println(dp[n-1][m-1])
}
