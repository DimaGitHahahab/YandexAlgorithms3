package YandexAlgorithms3

import (
	"fmt"
	"sort"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	var x int
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		a[i] = x
	}
	sort.Ints(a)
	dp := make([]int, n)
	dp[0] = 0
	switch n {
	case 1:
		fmt.Println(0)
	case 2:
		dp[1] = a[1] - a[0]
		fmt.Println(dp[1])
	default:
		dp[1] = a[1] - a[0]
		if n > 2 {
			dp[2] = a[2] - a[0]
			for i := 3; i < n; i++ {
				dp[i] = min(dp[i-1], dp[i-2]) + a[i] - a[i-1]
			}
		}
		fmt.Println(dp[n-1])
	}
}
