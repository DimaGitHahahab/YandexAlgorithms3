package YandexAlgorithms3

import "fmt"

func myMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var n int
	fmt.Scan(&n)
	first := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&first[i])
	}
	var m int
	fmt.Scan(&m)
	second := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&second[i])
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		if second[0] == first[i] {
			dp[i][0] = 1
		} else if i == 0 {
			dp[i][0] = 0
		} else {
			dp[i][0] = dp[i-1][0]
		}
	}
	for i := 0; i < m; i++ {
		if first[0] == second[i] {
			dp[0][i] = 1
		} else if i == 0 {
			dp[0][i] = 0
		} else {
			dp[0][i] = dp[0][i-1]
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if first[i] == second[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = myMax(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	ans := make([]int, 0)
	i, j := n-1, m-1
	for i >= 0 && j >= 0 {
		if first[i] == second[j] {
			ans = append(ans, first[i])
			i--
			j--
		} else {
			if i > 0 && j > 0 {
				if dp[i-1][j] > dp[i][j-1] {
					i--
				} else {
					j--
				}
			} else {
				if i > 0 {
					i--
				} else {
					j--
				}
			}
		}
	}
	for i := len(ans) - 1; i >= 0; i-- {
		fmt.Printf("%d ", ans[i])
	}
}
