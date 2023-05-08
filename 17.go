package YandexAlgorithms3

import "fmt"

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
	firstPlayerDeck, secondPlayerDeck := queue{}, queue{}
	var x int
	for i := 0; i < 5; i++ {
		fmt.Scan(&x)
		firstPlayerDeck.push(x)
	}
	for i := 0; i < 5; i++ {
		fmt.Scan(&x)
		secondPlayerDeck.push(x)
	}
	timer := 1000000
	for timer > 0 {
		if firstPlayerDeck.size() == 0 {
			fmt.Println("second", 1000000-timer)
			return
		}
		if secondPlayerDeck.size() == 0 {
			fmt.Println("first", 1000000-timer)
			return
		}
		firstPlayerCard := firstPlayerDeck.pop()
		secondPlayerCard := secondPlayerDeck.pop()
		if firstPlayerCard == 0 && secondPlayerCard == 9 {
			firstPlayerDeck.push(firstPlayerCard)
			firstPlayerDeck.push(secondPlayerCard)
		} else if firstPlayerCard == 9 && secondPlayerCard == 0 {
			secondPlayerDeck.push(firstPlayerCard)
			secondPlayerDeck.push(secondPlayerCard)
		} else if firstPlayerCard > secondPlayerCard {
			firstPlayerDeck.push(firstPlayerCard)
			firstPlayerDeck.push(secondPlayerCard)
		} else {
			secondPlayerDeck.push(firstPlayerCard)
			secondPlayerDeck.push(secondPlayerCard)
		}
		timer--
	}
	fmt.Println("botva")
	return
}
