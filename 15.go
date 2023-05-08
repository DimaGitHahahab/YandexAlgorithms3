package YandexAlgorithms3

import "fmt"

type Stack struct {
	data []int
}

func (s *Stack) Push(v int) {
	s.data = append(s.data, v)
}

func (s *Stack) Pop() int {
	x := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return x
}

func (s *Stack) Peek() int {
	return s.data[len(s.data)-1]
}

func (s *Stack) Size() int {
	return len(s.data)
}

func main() {
	houses, places := new(Stack), new(Stack)
	var n int
	fmt.Scan(&n)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = -1
	}
	var in int
	for i := 0; i < n; i++ {
		fmt.Scan(&in)
		if houses.Size() == 0 || houses.Peek() < in {
			houses.Push(in)
			places.Push(i)
		} else {
			for houses.Size() > 0 && houses.Peek() > in {
				ans[places.Pop()] = i
				houses.Pop()
			}
			houses.Push(in)
			places.Push(i)
		}
	}
	for i := 0; i < n; i++ {
		fmt.Print(ans[i], " ")
	}
}
