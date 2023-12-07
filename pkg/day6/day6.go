package day6

import (
	"fmt"
	"log"
	"math"
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

func parseInput(str string) ([]int, []int) {
	raceDuration := []int{}
	raceDistance := []int{}
	lines := strings.Split(str, "\n")
	durationStrs := strings.Fields(lines[0])
	for i := 1; i < len(durationStrs); i++ {
		n, err := strconv.Atoi(durationStrs[i])
		if err != nil {
			panic(err)
		}
		raceDuration = append(raceDuration, n)
	}
	distanceStrs := strings.Fields(lines[1])
	for i := 1; i < len(distanceStrs); i++ {
		n, err := strconv.Atoi(distanceStrs[i])
		if err != nil {
			panic(err)
		}
		raceDistance = append(raceDistance, n)
	}

	return raceDuration, raceDistance
}

func calcNumOfWaysToBeat(duration, distance int64) int64 {
	dur := float64(duration)
	dis := float64(distance)
	D := math.Pow(dur, 2) - 4*dis
	x1 := (-dur - math.Sqrt(D)) / -2
	x2 := (-dur + math.Sqrt(D)) / -2
	diff := x1 - math.Ceil(x2)
	if diff == math.Ceil(diff) {
		diff = math.Ceil(diff) - 1
	} else {
		diff = math.Ceil(diff)
	}
	return int64(diff)
}

func Solve(filename string) int64 {
	str := readFile(filename)
	raceDuration, raceDistance := parseInput(str)
	result := int64(1)
	for i := 0; i < len(raceDistance); i++ {
		n := calcNumOfWaysToBeat(int64(raceDuration[i]), int64(raceDistance[i]))
		fmt.Printf("%d) n = %d\n", i, n)
		result *= n
	}
	return result
}

func parseInput2(str string) (int64, int64) {
	lines := strings.Split(str, "\n")
	durationStrs := strings.Fields(lines[0])
	durationStr := ""
	for i := 1; i < len(durationStrs); i++ {
		durationStr += durationStrs[i]
	}
	distanceStrs := strings.Fields(lines[1])
	distanceStr := ""
	for i := 1; i < len(distanceStrs); i++ {
		distanceStr += distanceStrs[i]
	}
	duration, err := strconv.ParseInt(durationStr, 10, 64)
	if err != nil {
		panic(err)
	}
	distance, err := strconv.ParseInt(distanceStr, 10, 64)
	if err != nil {
		panic(err)
	}
	return duration, distance
}

func Solve2(filename string) int64 {
	str := readFile(filename)
	duration, distance := parseInput2(str)
	result := calcNumOfWaysToBeat(duration, distance)
	return result
}
