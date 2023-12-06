package day3

import (
	"fmt"
	"log"
	"os"
	"strconv"
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
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			grid[i][j] = filecontent[i*x+j+i]
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
	ok := false
	for i := 0; i < len(grid); i++ {
		fmt.Printf("%d)\t", i)
		for j := 0; j < len(grid[0]); j++ {
			if isNumber(grid[i][j]) {
				if !ok {
					ok = checkNumForAdjacentSymbols(grid, i, j)
				}
				bufn += string(grid[i][j])
			} else {
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
