package YandexAlgorithms3

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type queue struct {
	data []int
}

func (q *queue) push(i int) {
	q.data = append(q.data, i)
}

func (q *queue) pop() int {
	x := q.data[0]
	q.data = q.data[1:]
	return x
}

func (q *queue) front() int {
	return q.data[0]
}

func (q *queue) size() int {
	return len(q.data)
}

func (q *queue) clear() {
	q.data = q.data[:0]
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
	q := queue{}
	for _, line := range lines {
		tokens := strings.Fields(line)
		switch tokens[0] {
		case "push":
			v, _ := strconv.Atoi(tokens[1])
			q.push(v)
			output.WriteString("ok\n")
		case "pop":
			if q.size() == 0 {
				output.WriteString("error\n")
			} else {
				output.WriteString(strconv.Itoa(q.pop()) + "\n")
			}
		case "front":
			if q.size() == 0 {
				output.WriteString("error\n")
			} else {
				output.WriteString(strconv.Itoa(q.front()) + "\n")
			}
		case "size":
			output.WriteString(strconv.Itoa(q.size()) + "\n")
		case "clear":
			q.clear()
			output.WriteString("ok\n")
		case "exit":
			output.WriteString("bye\n")
			os.Exit(0)
		}
	}
}
