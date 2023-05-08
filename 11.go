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
	s := Stack{}
	for i := 0; i < len(lines); i++ {
		tokens := strings.Fields(lines[i])
		switch tokens[0] {
		case "push":
			v, _ := strconv.Atoi(tokens[1])
			s.push(v)
			output.WriteString("ok\n")

		case "pop":
			v, l := s.pop()
			if l == "ok" {
				vInt := strconv.Itoa(v)
				output.WriteString(vInt + "\n")
			} else {
				output.WriteString(l + "\n")
			}
		case "back":
			v, l := s.back()
			if l == "ok" {
				vInt := strconv.Itoa(v)
				output.WriteString(vInt + "\n")

			} else {
				output.WriteString(l + "\n")
			}
		case "size":
			vInt := strconv.Itoa(s.size())
			output.WriteString(vInt + "\n")
		case "clear":
			s.clear()
			output.WriteString("ok" + "\n")
		case "exit":
			output.WriteString("bye")
			os.Exit(0)
		}
	}

}
