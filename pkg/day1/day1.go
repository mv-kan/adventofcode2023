package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func isNumber(r byte) bool {
	return (r-'0' < 10 && r-'0' > 0)
}

// https://adventofcode.com/2023/day/1
func Solve(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := 0
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		str := scanner.Text()
		strN := ""
		for i := 0; i < len(str); i++ {
			if isNumber(str[i]) {
				strN += string(str[i])
				break
			}
		}
		for i := len(str) - 1; i >= 0; i-- {
			if isNumber(str[i]) {
				strN += string(str[i])
				break
			}
		}
		num, err := strconv.Atoi(strN)
		if err != nil {
			panic(err)
		}
		result += num
	}
	return result
}

var nums = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var numsReverse = []string{"orez", "eno", "owt", "eerht", "ruof", "evif", "xis", "neves", "thgie", "enin"}
var str2num = map[string]string{
	"zero":  "0",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}
var str2numRev = map[string]string{
	"orez":  "0",
	"eno":   "1",
	"owt":   "2",
	"eerht": "3",
	"ruof":  "4",
	"evif":  "5",
	"xis":   "6",
	"neves": "7",
	"thgie": "8",
	"enin":  "9",
}

func checkStr(str string, reversed bool) (string, bool) {
	checkFor := 0
	passedChars := 0
	nums_ := nums
	if reversed {
		nums_ = numsReverse
	}
	l := len(nums_[checkFor])
	for i := 0; i < l && i < len(str); i++ {
		if checkFor > 9 {
			return "", false
		}
		if str[i] != nums_[checkFor][i] {
			checkFor++
			if checkFor <= 9 {
				l = len(nums_[checkFor])
			}
			i = -1 // because i++
			passedChars = 0
		} else {
			passedChars++
		}
	}
	return str, true
}

// function, which takes a string as
// argument and return the reverse of string.
func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}

func getFirstNumFromStr(str string, reversed bool) string {
	buf := ""
	strN := ""
	str2num_ := str2num
	if reversed {
		str2num_ = str2numRev
	}
	for i := 0; i < len(str); i++ {
		if isNumber(str[i]) {
			strN += string(str[i])
			buf = ""
			break
		} else {
			buf += string(str[i])
			_, ok := checkStr(buf, reversed)
			if ok {
				n, ok := str2num_[buf]
				if ok {
					strN += n
					buf = ""
					break
				}
			} else {
				if len(buf) >= 3 {
					buf = string(buf[1:])
				} else {
					buf = string(buf[len(buf)-1])
				}
			}
		}
	}
	return strN
}

func SolvePart2(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := 0

	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		str := scanner.Text()
		n1 := getFirstNumFromStr(str, false)
		str = reverse(str)
		n2 := getFirstNumFromStr(str, true)

		num, err := strconv.Atoi(n1 + n2)
		if err != nil {
			panic(err)
		}
		result += num
	}
	return result
}
