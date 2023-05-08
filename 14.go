package YandexAlgorithms3

import (
	"fmt"
	"math"
)

type Stack struct {
	data []int
}

func (s *Stack) push(x int) {
	s.data = append(s.data, x)
	return
}

func (s *Stack) pop() int {
	x := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return x
}

func (s *Stack) back() int {
	return s.data[len(s.data)-1]

}

func (s *Stack) size() int {
	return len(s.data)
}

func (s *Stack) clear() {
	s.data = s.data[:0]
}

func main() {
	var n, y int
	last := math.MinInt32
	fmt.Scan(&n)
	a := make([]int, n, n)
	enter := Stack{data: a}
	for i := n - 1; i >= 0; i-- {
		fmt.Scan(&y)
		enter.data[i] = y
	}
	wall := Stack{}

	for enter.size() > 0 {
		if wall.size() == 0 || enter.back() < wall.back() {
			wall.push(enter.pop())
		} else {
			current := wall.pop()
			if current < last {
				fmt.Print("NO")
				return
			} else {
				last = current
			}
		}
	}
	for wall.size() > 0 {
		current := wall.pop()
		if current < last {
			fmt.Print("NO")
			return
		} else {
			last = current
		}
	}
	fmt.Print("YES")
}
