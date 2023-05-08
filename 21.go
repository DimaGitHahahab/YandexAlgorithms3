package YandexAlgorithms3

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n+1)
	if n == 1 {
		fmt.Println(2)
		return
	}
	if n == 2 {
		fmt.Println(4)
		return
	}
	if n == 3 {
		fmt.Println(7)
		return
	}
	a[0] = 1
	a[1] = 2
	a[2] = 4

	for i := 3; i <= n; i++ {
		a[i] = a[i-1] + a[i-2] + a[i-3]
	}
	fmt.Println(a[n])
}
