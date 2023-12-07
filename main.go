package main

import (
	"fmt"

	"github.com/mv-kan/adventofcode2023/pkg/day6"
)

func main() {
	result := day6.Solve("./pkg/day6/input_test.txt")
	fmt.Printf("result = %d\n", result)
	result = day6.Solve2("./pkg/day6/input.txt")
	fmt.Printf("result part 2 = %d\n", result)
}
