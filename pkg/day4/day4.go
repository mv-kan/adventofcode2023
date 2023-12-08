package day4

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func parseCardId(str string) (int, string) {
	s := strings.Split(str, ":")
	gameStr := strings.Fields(s[0])
	id, err := strconv.Atoi(gameStr[1])
	if err != nil {
		panic(err)
	}
	return id, s[1]
}

func splitIntoWinning(str string) []string {
	return strings.Split(str, "|")
}

func trimSplitAndCompare(cards []string) (points int) {
	winningCards := map[int]bool{}
	winningCardsStr := strings.Trim(cards[0], " ")
	winningCardsStrSl := strings.Fields(winningCardsStr)
	for i := 0; i < len(winningCardsStrSl); i++ {
		n, err := strconv.Atoi(winningCardsStrSl[i])
		if err != nil {
			panic(err)
		}
		winningCards[n] = true
	}
	ourCardsStr := strings.Trim(cards[1], " ")
	ourCardStrSl := strings.Fields(ourCardsStr)
	ourCards := []int{}
	for i := 0; i < len(ourCardStrSl); i++ {
		if ourCardStrSl[i] == "" {
			continue
		}
		n, err := strconv.Atoi(ourCardStrSl[i])
		if err != nil {
			panic(err)
		}
		ourCards = append(ourCards, n)
	}

	for i := 0; i < len(ourCards); i++ {
		_, ok := winningCards[ourCards[i]]
		if ok {
			points++
		}
	}
	return
}
func readFile(filename string) string {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(file)
}
func Solve(filename string) int {
	filecontent := readFile(filename)
	cards := strings.Split(filecontent, "\n")
	result := 0
	cardsNum := len(cards)
	for i := 0; i < cardsNum; i++ {
		str := cards[i]
		id, str := parseCardId(str)
		winOurCards := splitIntoWinning(str)
		points := trimSplitAndCompare(winOurCards)
		for i := 1; i <= points; i++ {
			cards = append(cards, cards[id+i-1])
		}
		cardsNum = len(cards)
		result++
	}
	return result
}
