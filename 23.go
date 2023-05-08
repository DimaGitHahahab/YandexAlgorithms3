package YandexAlgorithms3

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n+1)
	a[1] = 0
	for i := 2; i <= n; i++ {
		a[i] = a[i-1]
		if i%2 == 0 {
			if a[i/2] < a[i] {
				a[i] = a[i/2]
			}
		}
		if i%3 == 0 {
			if a[i/3] < a[i] {
				a[i] = a[i/3]
			}
		}
		a[i]++
	}
	fmt.Println(a[n])
	steps := make([]int, 0)
	steps = append(steps, n)
	for n > 1 {
		if n%3 == 0 && a[n/3] == a[n]-1 {
			n /= 3
		} else if n%2 == 0 && a[n/2] == a[n]-1 {
			n /= 2
		} else {
			n--
		}
		steps = append(steps, n)
	}
	for i := len(steps) - 1; i >= 0; i-- {
		fmt.Print(steps[i], " ")
	}
}
