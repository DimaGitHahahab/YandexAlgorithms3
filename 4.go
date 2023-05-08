package YandexAlgorithms3

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	var row, col int
	fmt.Scan(&row, &col)
	var upRow, upCol, downRow, downCol int
	upRow = row - k/2
	upCol = col - k%2
	if upCol == 0 {
		upRow--
		upCol = 2
	}
	downRow = row + k/2
	downCol = col + k%2
	if downCol == 3 {
		downRow++
		downCol = 1
	}
	var amountOfRows int
	if n%2 == 0 {
		amountOfRows = n / 2
	} else {
		amountOfRows = n/2 + 1
	}
	if upRow < 1 && downRow > amountOfRows {
		fmt.Println(-1)
		return
	}
	if upRow < 1 {
		if downRow == amountOfRows && downCol == 2 && n%2 == 1 {
			fmt.Println(-1)
			return
		} else {
			fmt.Println(downRow, downCol)
			return
		}
	}
	if downRow > amountOfRows {
		fmt.Println(upRow, upCol)
		return
	}
	upDiff := row - upRow
	downDiff := downRow - row
	if upDiff < downDiff {
		fmt.Println(upRow, upCol)
		return
	} else {
		if downRow == amountOfRows && downCol == 2 && n%2 == 1 {
			fmt.Println(upRow, upCol)
			return
		}
		fmt.Println(downRow, downCol)
		return
	}
}
