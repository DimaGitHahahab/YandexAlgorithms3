package YandexAlgorithms3

import "fmt"

type heap struct {
	data []int
}

func (h *heap) insert(n int) {
	h.data = append(h.data, n)
	i := len(h.data) - 1
	for i > 0 && h.data[i] < h.data[(i-1)/2] {
		h.data[(i-1)/2], h.data[i] = h.data[i], h.data[(i-1)/2]
		i = (i - 1) / 2
	}
}

func (h *heap) removeMin() (x int) {
	h.data[0], h.data[len(h.data)-1] = h.data[len(h.data)-1], h.data[0]
	x = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	i := 0
	for 2*i+1 < len(h.data) {
		j := 2*i + 1
		if 2*i+2 < len(h.data) && h.data[2*i+2] < h.data[j] {
			j = 2*i + 2
		}
		if h.data[i] <= h.data[j] {
			break
		} else {
			h.data[i], h.data[j] = h.data[j], h.data[i]
			i = j
		}
	}
	return
}

func main() {
	a := make([]int, 0)
	var n int
	fmt.Scan(&n)
	var y int
	for i := 0; i < n; i++ {
		fmt.Scan(&y)
		a = append(a, y)
	}
	h := new(heap)
	for i := 0; i < len(a); i++ {
		h.insert(a[i])
	}
	for i := 0; i < len(a); i++ {
		a[i] = h.removeMin()
	}
	for _, v := range a {
		fmt.Print(v, " ")
	}
}
