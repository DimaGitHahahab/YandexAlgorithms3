package YandexAlgorithms3

import "fmt"

type heap struct {
	data []int
}

func (h *heap) insert(n int) {
	h.data = append(h.data, n)
	i := len(h.data) - 1
	for i > 0 && h.data[i] > h.data[(i-1)/2] {
		h.data[(i-1)/2], h.data[i] = h.data[i], h.data[(i-1)/2]
		i = (i - 1) / 2
	}
}

func (h *heap) removeMax() (x int) {
	h.data[0], h.data[len(h.data)-1] = h.data[len(h.data)-1], h.data[0]
	x = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	i := 0
	for 2*i+1 < len(h.data) {
		j := 2*i + 1
		if 2*i+2 < len(h.data) && h.data[2*i+2] > h.data[j] {
			j = 2*i + 2
		}
		if h.data[i] >= h.data[j] {
			break
		} else {
			h.data[i], h.data[j] = h.data[j], h.data[i]
			i = j
		}
	}
	return
}

func main() {
	h := new(heap)
	var n int
	fmt.Scan(&n)
	var y, z int
	for i := 0; i < n; i++ {
		fmt.Scan(&y)
		switch y {
		case 0:
			fmt.Scan(&z)
			h.insert(z)
		case 1:
			fmt.Println(h.removeMax())
		}
	}
}
