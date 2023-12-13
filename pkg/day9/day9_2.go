package day9

import "fmt"

func calcPrev(nums []int) int {
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
	toSubtract := 0
	for i := len(diffss) - 1; i >= 0; i-- {
		toSubtract = diffss[i][0] - toSubtract
	}
	result = toSubtract
	return result
}

func Solve2(filename string) int {
	nums := parseInput(filename)
	result := 0
	for i := 0; i < len(nums); i++ {
		result += calcPrev(nums[i])
		fmt.Printf("r = %d\n", result)
	}
	return result
}
