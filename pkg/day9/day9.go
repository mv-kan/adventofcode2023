package day9

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile(filename string) string {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(file)
}

func parseInput(filename string) (result [][]int) {
	filecontent := readFile(filename)
	s := strings.Split(filecontent, "\n")
	result = [][]int{}
	for i := 0; i < len(s); i++ {
		row := []int{}
		nums := strings.Fields(s[i])
		for j := 0; j < len(nums); j++ {
			n, err := strconv.Atoi(nums[j])
			if err != nil {
				panic(err)
			}
			row = append(row, n)
		}
		result = append(result, row)
	}
	return
}

func calcNext(nums []int) int {
	result := 0
	diffss := [][]int{nums}
	diffssLen := 1
	for i := 0; i < diffssLen; i++ {
		allZero := true
		diffs := []int{}
		for j := 1; j < len(diffss[i]); j++ {
			diff := diffss[i][j] - diffss[i][j-1]
			diffs = append(diffs, diff)
			if allZero && diff != 0 {
				allZero = false
			}
		}
		diffss = append(diffss, diffs)
		diffssLen = len(diffss)
		if allZero {
			break
		}
	}
	toAdd := 0
	for i := len(diffss) - 1; i >= 0; i-- {
		toAdd = diffss[i][len(diffss[i])-1] + toAdd
	}
	result = toAdd
	return result
}

func Solve(filename string) int {
	nums := parseInput(filename)
	result := 0
	for i := 0; i < len(nums); i++ {
		result += calcNext(nums[i])
		fmt.Printf("r = %d\n", result)
	}
	return result
}
