package main

import (
	"fmt"

	"github.com/mv-kan/adventofcode2023/pkg/day2"
)

func main() {
	// result := day1.Solve("./pkg/day1/input.txt")
	// fmt.Printf("Day1: %d\n", result)
	// result = day1.SolvePart2("./pkg/day1/input_part2.txt")
	// // result = day1.SolvePart2("./pkg/day1/input_test.txt")
	// fmt.Printf("Day1 part2: %d\n", result)
	result := day2.Solve("./pkg/day2/input_part1.txt")
	fmt.Printf("Day2: %d\n", result)
	result = day2.Solve2("./pkg/day2/input_part2.txt")
	fmt.Printf("Day2 part2: %d\n", result)
}
