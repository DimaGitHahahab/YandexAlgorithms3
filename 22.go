package YandexAlgorithms3

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	steps := make([]int, n+1)
	steps[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= k; j++ {
			if i-j >= 1 {
				steps[i] += steps[i-j]
			}
		}
	}
	fmt.Println(steps[n])
}
