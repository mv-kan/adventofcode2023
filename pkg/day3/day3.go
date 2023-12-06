package day3

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isNumber(r byte) bool {
	return (r >= 48 && r <= 57)
}

func readFile(filename string) string {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(file)
}

// 1187317
// 1277016
func getDims(input string) (i int, j int) {
	count := 0
	for index, char := range input {
		if char == '\n' {
			count++
			if j == 0 {
				j = index
			}
		}
	}
	i = count + 1
	return
}

func fillGrid(filecontent string, grid [][]byte) {
	x := len(grid)
	y := len(grid[0])
	str := strings.Split(filecontent, "\n")
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			grid[i][j] = str[i][j]
		}
	}
}

func checkNumForAdjacentSymbols(grid [][]byte, celli, cellj int) bool {
	for i := celli - 1; i <= celli+1; i++ {
		for j := cellj - 1; j <= cellj+1; j++ {
			ok := checkCellForSymbol(grid, celli, cellj, i, j)
			if ok {
				return true
			}
		}
	}
	return false
}

func checkCellForSymbol(grid [][]byte, i, j int, symboli, symbolj int) bool {
	if symboli >= len(grid) || symboli < 0 {
		return false
	}
	if symbolj >= len(grid[0]) || symbolj < 0 {
		return false
	}
	if !isNumber(grid[i][j]) {
		panic("Lol this should not happen")
	}
	return grid[symboli][symbolj] != '.' && !isNumber(grid[symboli][symbolj])
}

func parseNumberFromCoords(grid [][]byte, c []coords) (int, int, bool) {
	n1, n2 := 0, 0
	n1BeginCoords := coords{}
	n2BeginCoords := coords{}

	getBeginOfNumber := func(co coords) coords {
		result := coords{
			row: co.row,
			col: co.col,
		}
		for i := co.col; i >= 0; i-- {
			if !isNumber(grid[co.row][i]) {
				result.col = i + 1
				break
			} else {
				result.col = i
			}
		}

		return result
	}
	parseNumberFromBegin := func(co coords) int {
		str := ""
		for i := co.col; i < len(grid[0]); i++ {
			if !isNumber(grid[co.row][i]) {
				break
			}
			str += string(grid[co.row][i])
		}
		n, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		return n
	}
	for i := 0; i < len(c); i++ {
		beginCoords := getBeginOfNumber(c[i])
		n := parseNumberFromBegin(beginCoords)
		if n1 == 0 || beginCoords == n1BeginCoords {
			n1 = n
			n1BeginCoords = beginCoords
		} else if n2 == 0 || beginCoords == n2BeginCoords {
			n2 = n
			n2BeginCoords = beginCoords
		}
	}
	if n1 == 0 || n2 == 0 {
		return 0, 0, false
	}
	return n1, n2, true
}

type coords struct {
	row, col int
}

func checkGear(grid [][]byte, celli, cellj int) []coords {
	result := []coords{}
	for i := celli - 1; i <= celli+1; i++ {
		for j := cellj - 1; j <= cellj+1; j++ {
			if i >= len(grid) || i < 0 {
				continue
			}
			if j >= len(grid[0]) || j < 0 {
				continue
			}
			if isNumber(grid[i][j]) {
				result = append(result, coords{
					row: i,
					col: j,
				})
			}
		}
	}
	return result
}

func Solve(filename string) int {
	filecontent := readFile(filename)
	x, y := getDims(filecontent)
	grid := make([][]byte, x)
	for i := 0; i < x; i++ {
		grid[i] = make([]byte, y)
	}
	fillGrid(filecontent, grid)

	result := 0
	bufn := ""
	for i := 0; i < len(grid); i++ {
		fmt.Printf("%d)\t", i)
		ok := false
		for j := 0; j < len(grid[0]); j++ {
			if isNumber(grid[i][j]) {
				if !ok {
					ok = checkNumForAdjacentSymbols(grid, i, j)
				}
				bufn += string(grid[i][j])
			}
			if j == len(grid[0])-1 || !isNumber(grid[i][j]) {
				if ok {
					n, err := strconv.Atoi(bufn)
					if err != nil {
						panic(err)
					}
					fmt.Printf("%d  ", n)
					result += n
				}
				bufn = ""
				ok = false
			}
			// fmt.Printf("%s", string(grid[i][j]))
		}
		fmt.Println()
	}

	return result
}

func Solve2(filename string) int {
	filecontent := readFile(filename)
	x, y := getDims(filecontent)
	grid := make([][]byte, x)
	for i := 0; i < x; i++ {
		grid[i] = make([]byte, y)
	}
	fillGrid(filecontent, grid)

	result := 0
	bufn := ""
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '*' {
				neighborNumsCoords := checkGear(grid, i, j)
				n1, n2, ok := parseNumberFromCoords(grid, neighborNumsCoords)
				if !ok {
					continue
				}
				bufn += string(grid[i][j])
				result += n1 * n2
			}
		}
	}

	return result
}
