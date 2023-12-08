package day7

type rank int

const (
	fiveOfKind rank = iota
	fourOfKind
	fullHouse
	threeOfKind
	twoPair
	onePair
	highCard
)

func evaluateRank(hand string) rank {
	pairs := 0
	triplet := 0
	quartet := 0
	quintet := 0
	for i := 0; i < len(hand); i++ {
		similar := 0
		for j := i + 1; j < len(hand); j++ {
			if hand[i] == hand[j] {
				similar++
			}
		}
		if similar == 4 {
			quintet++
			break
		} else if similar == 3 {
			quartet++
			break
		} else if similar == 2 {
			triplet++
		} else if similar == 1 {
			pairs++
		}
	}
	if quintet == 1 {
		return fiveOfKind
	} else if quartet == 1 {
		return fourOfKind
	} else if triplet == 1 && pairs == 1 {
		return fullHouse
	} else if triplet == 1 {
		return threeOfKind
	} else if pairs == 2 {
		return twoPair
	} else if pairs == 1 {
		return onePair
	} else {
		return highCard
	}
}

func evaluateSubrank(hand string) int {

}
