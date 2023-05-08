package YandexAlgorithms3

import (
	"fmt"
	"sort"
)

func binarySearch(arr *[]int, target int) int {
	l, r := -1, len(*arr)
	for r > l+1 {
		m := (l + r) / 2
		if (*arr)[m] <= target {
			l = m
		} else {
			r = m
		}
	}
	return l
}

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	sort.Ints(arr)
	uniqueArray := make([]int, 0)
	for i := 0; i < n; i++ {
		if i == 0 || arr[i] != arr[i-1] {
			uniqueArray = append(uniqueArray, arr[i])
		}
	}
	var k int
	fmt.Scan(&k)
	q := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Scan(&q[i])
	}
	for i := 0; i < k; i++ {
		count := 0
		q[i]--
		res := binarySearch(&uniqueArray, q[i])
		if res != -1 {
			count = res + 1
		}
		fmt.Println(count)
	}
}
