package YandexAlgorithms3

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strings"
)

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
	m := make(map[string]int)
	var mx int
	for _, line := range lines {
		for _, letter := range line {
			if letter != ' ' {
				m[string(letter)]++
				if m[string(letter)] > mx {
					mx = m[string(letter)]
				}
			}
		}
	}
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	depth := mx
	for i := 0; i < depth; i++ {
		for _, k := range keys {
			if m[k] == mx {
				output.WriteString("#")
				m[k]--
			} else {
				output.WriteString(" ")
			}
		}
		mx--
		output.WriteString("\n")
	}
	for _, k := range keys {
		output.WriteString(k)
	}
}
