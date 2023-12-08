package main

import (
	"fmt"

	"github.com/mv-kan/adventofcode2023/pkg/day5"
)

func main() {
	result := day5.Solve2("./pkg/day5/input_test.txt")
	fmt.Println(result)
	result = day5.Solve2("./pkg/day5/input.txt")
	fmt.Println(result)
}
