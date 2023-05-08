package YandexAlgorithms3

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	letters := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&letters[i])
	}
	var ans int
	for i := 0; i < n-1; i++ {
		if letters[i]-letters[i+1] > 0 {
			ans += letters[i+1]
		} else {
			ans += letters[i]
		}
	}
	fmt.Println(ans)
}
