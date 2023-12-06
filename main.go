package main

import (
	"fmt"

	"github.com/mv-kan/adventofcode2023/pkg/day1"
)

func main() {
	result := day1.Solve("./pkg/day1/input.txt")
	fmt.Printf("Day1: %d\n", result)
	result = day1.SolvePart2("./pkg/day1/input_part2.txt")
	// result = day1.SolvePart2("./pkg/day1/input_test.txt")
	fmt.Printf("Day1 part2: %d\n", result)
}
