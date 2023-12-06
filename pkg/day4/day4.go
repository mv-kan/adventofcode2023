package day4

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseCardId(str string) (int, string) {
	s := strings.Split(str, ":")
	gameStr := strings.Split(s[0], " ")
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
	winningCardsStrSl := strings.Split(winningCardsStr, " ")
	for i := 0; i < len(winningCardsStrSl); i++ {
		n, err := strconv.Atoi(winningCardsStrSl[i])
		if err != nil {
			panic(err)
		}
		winningCards[n] = true
	}
	ourCardsStr := strings.Trim(cards[1], " ")
	ourCardStrSl := strings.Split(ourCardsStr, " ")
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
			if points == 0 {
				points = 1
			} else {
				points *= 2
			}
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
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	result := 0
	for scanner.Scan() {
		str := scanner.Text()
		_, str = parseCardId(str)
		cards := splitIntoWinning(str)
		result += trimSplitAndCompare(cards)
	}
	return result
}
