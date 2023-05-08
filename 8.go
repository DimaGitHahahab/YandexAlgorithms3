package YandexAlgorithms3

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var minX, maxX, minY, maxY, x, y int
	fmt.Scan(&x, &y)
	minX, minY, maxX, maxY = x, y, x, y
	for i := 0; i < n-1; i++ {
		fmt.Scan(&x, &y)
		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
	}
	fmt.Println(minX, minY, maxX, maxY)
}
