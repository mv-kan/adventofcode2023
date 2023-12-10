package day7

import (
	"bufio"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type rank int

const (
	only int = 1 << iota
	pair
	twoPairs
	triplet
	quartet
	quintet
)

const (
	fiveOfKind  rank = rank(quintet)
	fourOfKind  rank = rank(quartet)
	fullHouse   rank = rank(triplet | pair)
	threeOfKind rank = rank(triplet)
	twoPair     rank = rank(twoPairs) // XOR
	onePair     rank = rank(pair)
	highCard    rank = rank(only)
)

func evaluateRank(hand string) rank {
	seenCards := ""
	card2num := map[byte]int{}
	for i := 0; i < len(hand); i++ {
		n, ok := card2num[hand[i]]
		if !ok {
			seenCards += string(hand[i])
			card2num[hand[i]] = 0
			n = 0
		}
		card2num[hand[i]] = n + 1
	}
	handRank := highCard
	for i := 0; i < len(seenCards); i++ {
		n, ok := card2num[seenCards[i]]
		if !ok {
			panic(ok)
		}
		switch n {
		case 2:
			if handRank == threeOfKind {
				handRank = fullHouse
			} else if handRank == onePair {
				handRank = twoPair
			} else {
				handRank = onePair
			}
		case 3:
			if handRank == onePair {
				handRank = fullHouse
			} else {
				handRank = threeOfKind
			}
		case 4:
			handRank = fourOfKind
		case 5:
			handRank = fiveOfKind
		}
	}
	return handRank
}

func evaluateSubrank(hand string) int {
	card2dec := map[byte]int{
		'A': 13,
		'K': 12,
		'Q': 11,
		'J': 10,
		'T': 9,
		'9': 8,
		'8': 7,
		'7': 6,
		'6': 5,
		'5': 4,
		'4': 3,
		'3': 2,
		'2': 1,
	}
	// 13 possible cards
	subrank := 0
	for i := 0; i < len(hand); i++ {
		place := int(math.Pow(13, float64(len(hand)-i-1)))
		decimal, ok := card2dec[hand[i]]
		if !ok {
			panic(ok)
		}
		subrank += place * decimal
	}

	return subrank
}

type handBet struct {
	hand string
	bet  int
}

func parseInput(filename string) (hands []handBet) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		str := scanner.Text()
		s := strings.Fields(str)
		bet, err := strconv.Atoi(s[1])
		if err != nil {
			panic(err)
		}
		hands = append(hands, handBet{
			hand: s[0],
			bet:  bet,
		})
	}
	return
}

func Solve(filename string) int {
	hands := parseInput(filename)
	sort.SliceStable(hands, func(i, j int) bool {
		return evaluateSubrank(hands[i].hand) < evaluateSubrank(hands[j].hand)
	})
	sort.SliceStable(hands, func(i, j int) bool {
		return evaluateRank(hands[i].hand) < evaluateRank(hands[j].hand)
	})
	result := 0
	for i := 0; i < len(hands); i++ {
		result += (i + 1) * hands[i].bet
	}
	return result
}
