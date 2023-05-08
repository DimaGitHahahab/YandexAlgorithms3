package YandexAlgorithms3

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	data []int
}

func (s *Stack) push(x int) {
	s.data = append(s.data, x)
	return
}

func (s *Stack) pop() (int, string) {
	if len(s.data) > 0 {
		x := s.data[len(s.data)-1]
		s.data = s.data[:len(s.data)-1]
		return x, "ok"
	} else {
		return -1, "error"
	}
}

func (s *Stack) back() (int, string) {
	if len(s.data) > 0 {
		x := s.data[len(s.data)-1]
		return x, "ok"
	} else {
		return -1, "error"
	}
}

func (s *Stack) size() int {
	return len(s.data)
}

func (s *Stack) clear() {
	s.data = s.data[:0]
}

func main() {
	var lines []string
	file, _ := os.Open("input.txt")
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}
		lines = append(lines, strings.TrimSpace(line))
		if err != nil {
			break
		}
	}
	output, _ := os.Create("output.txt")
	defer output.Close()
	stack := Stack{}
	buffer := strings.Split(lines[0], " ")
	for _, symbol := range buffer {
		if symbol != "*" && symbol != "+" && symbol != "-" {
			number, _ := strconv.Atoi(symbol)
			stack.push(number)
		} else {
			firstNumber, _ := stack.pop()
			secondNumber, _ := stack.pop()
			var result int
			switch symbol {
			case "*":
				result = firstNumber * secondNumber
			case "+":
				result = firstNumber + secondNumber
			case "-":
				result = secondNumber - firstNumber
			}
			stack.push(result)
		}
	}
	result, _ := stack.back()
	output.WriteString(strconv.Itoa(result))
}
