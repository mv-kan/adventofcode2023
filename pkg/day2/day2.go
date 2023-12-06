package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	redCubes   = 12
	greenCubes = 13
	blueCubes  = 14
)

func parseGameId(str string) (int, string) {
	s := strings.Split(str, ":")
	gameStr := strings.Split(s[0], " ")
	id, err := strconv.Atoi(gameStr[1])
	if err != nil {
		panic(err)
	}
	return id, s[1]
}

func splitIntoRounds(str string) []string {
	return strings.Split(str, ";")
}

func parseCubesAmount(str string) (red, green, blue int) {
	cubes := strings.Split(str, ",")
	for i := 0; i < len(cubes); i++ {
		n, cube := parseCube(cubes[i])
		switch cube {
		case "red":
			red = n
		case "green":
			green = n
		case "blue":
			blue = n
		default:
			panic(cube)
		}
	}
	return
}

func parseCube(str string) (int, string) {
	str = strings.TrimSpace(str)
	s := strings.Split(str, " ")
	n, err := strconv.Atoi(s[0])
	if err != nil {
		panic(err)
	}
	return n, s[1]
}

// https://adventofcode.com/2023/day/2
func Solve(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	result := 0
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		str := scanner.Text()
		ok := true
		id, str := parseGameId(str)
		fmt.Printf("ID: %d, str: %s\n", id, str)
		rounds := splitIntoRounds(str)
		for i := 0; i < len(rounds); i++ {
			red, green, blue := parseCubesAmount(rounds[i])
			if red > redCubes {
				ok = false
			}
			if green > greenCubes {
				ok = false
			}
			if blue > blueCubes {
				ok = false
			}
		}
		if ok {
			result += id
		}
	}
	return result
}

func Solve2(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	result := 0
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		str := scanner.Text()
		maxRed := 0
		maxGreen := 0
		maxBlue := 0
		id, str := parseGameId(str)
		fmt.Printf("ID: %d, str: %s\n", id, str)
		rounds := splitIntoRounds(str)
		for i := 0; i < len(rounds); i++ {
			red, green, blue := parseCubesAmount(rounds[i])
			if red > maxRed {
				maxRed = red
			}
			if green > maxGreen {
				maxGreen = green
			}
			if blue > maxBlue {
				maxBlue = blue
			}
		}
		result += maxRed * maxGreen * maxBlue
	}
	return result
}
