package YandexAlgorithms3

import (
	"fmt"
)

func myMin(nums ...int) int {
	min := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min
}

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	b := make([]int, n)
	c := make([]int, n)
	var x, y, z int
	for i := 0; i < n; i++ {
		fmt.Scan(&x, &y, &z)
		a[i], b[i], c[i] = x, y, z
	}
	dp := make([]int, n)
	switch n {
	case 1:
		dp[0] = a[0]
		fmt.Println(dp[0])
		return
	case 2:
		dp[0] = a[0]
		dp[1] = myMin(a[0]+a[1], b[0])
		fmt.Println(dp[1])
		return
	case 3:
		dp[0] = a[0]
		dp[1] = myMin(a[0]+a[1], b[0])
		dp[2] = myMin(a[0]+a[1]+a[2], b[0]+a[2], a[0]+b[1], c[0])
		fmt.Println(dp[2])
		return
	default:
		dp[0] = a[0]
		dp[1] = myMin(a[0]+a[1], b[0])
		dp[2] = myMin(a[0]+a[1]+a[2], b[0]+a[2], a[0]+b[1], c[0])
		for i := 3; i < n; i++ {
			dp[i] = myMin(dp[i-1]+a[i], dp[i-2]+b[i-1], dp[i-3]+c[i-2])
		}
		fmt.Println(dp[n-1])
	}
}
