package YandexAlgorithms3

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type Stack struct {
	data []string
}

func (s *Stack) push(x string) {
	s.data = append(s.data, x)
	return
}

func (s *Stack) pop() (string, string) {
	if len(s.data) > 0 {
		x := s.data[len(s.data)-1]
		s.data = s.data[:len(s.data)-1]
		return x, "ok"
	} else {
		return "", "error"
	}
}

func (s *Stack) back() (string, string) {
	if len(s.data) > 0 {
		x := s.data[len(s.data)-1]
		return x, "ok"
	} else {
		return "", "error"
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
	for i := 0; i < len(lines[0]); i++ {
		letter := string(lines[0][i])
		switch letter {
		case "(":
			stack.push(")")
		case "{":
			stack.push("}")
		case "[":
			stack.push("]")
		default:
			v, _ := stack.back()
			if stack.size() == 0 || v != letter {
				output.WriteString("no")
				return
			}
			stack.pop()
		}
	}
	if stack.size() == 0 {
		output.WriteString("yes")
	} else {
		output.WriteString("no")
	}
}
