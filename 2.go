package YandexAlgorithms3

import "fmt"

func main() {
	var k int
	var s string
	fmt.Scan(&k, &s)
	alf := make(map[string]int)
	for i := 0; i < len(s); i++ {
		alf[string(s[i])]++
	}
	var ans int
	for letter := range alf {
		i, tempK, j := 0, k, 0
		var max int
		for i < len(s) {
			for tempK > 0 && j < len(s) || (j < len(s) && string(s[j]) == letter && tempK == 0) {
				if string(s[j]) != letter {
					tempK--
				}
				max++
				j++
			}
			if max > ans {
				ans = max
			}
			if string(s[i]) != letter {
				tempK++
			}
			max--
			i++
		}
	}
	fmt.Println(ans)
}
