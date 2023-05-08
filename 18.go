package YandexAlgorithms3

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type deque struct {
	data []int
}

func (d *deque) push_front(n int) {
	d.data = append([]int{n}, d.data...)
}

func (d *deque) push_back(n int) {
	d.data = append(d.data, n)
}

func (d *deque) pop_front() (x int) {
	x = d.data[0]
	d.data = d.data[1:]
	return
}

func (d *deque) pop_back() (x int) {
	x = d.data[len(d.data)-1]
	d.data = d.data[:len(d.data)-1]
	return
}

func (d *deque) front() int {
	return d.data[0]
}

func (d *deque) back() int {
	return d.data[len(d.data)-1]
}

func (d *deque) size() int {
	return len(d.data)
}

func (d *deque) clear() {
	d.data = d.data[:0]
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
	d := new(deque)
	for _, line := range lines {
		tokens := strings.Split(line, " ")
		switch tokens[0] {
		case "push_front":
			v, _ := strconv.Atoi(tokens[1])
			d.push_front(v)
			output.WriteString("ok\n")
		case "push_back":
			v, _ := strconv.Atoi(tokens[1])
			d.push_back(v)
			output.WriteString("ok\n")
		case "pop_front":
			if d.size() == 0 {
				output.WriteString("error\n")
			} else {
				output.WriteString(strconv.Itoa(d.pop_front()) + "\n")
			}
		case "pop_back":
			if d.size() == 0 {
				output.WriteString("error\n")
			} else {
				output.WriteString(strconv.Itoa(d.pop_back()) + "\n")
			}
		case "front":
			if d.size() == 0 {
				output.WriteString("error\n")
			} else {
				output.WriteString(strconv.Itoa(d.front()) + "\n")
			}
		case "back":
			if d.size() == 0 {
				output.WriteString("error\n")
			} else {
				output.WriteString(strconv.Itoa(d.back()) + "\n")
			}
		case "clear":
			d.clear()
			output.WriteString("ok\n")
		case "size":
			output.WriteString(strconv.Itoa(d.size()) + "\n")
		case "exit":
			output.WriteString("bye\n")
			os.Exit(0)
		}
	}
}
