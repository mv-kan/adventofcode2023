package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type internal int64

func parseInput(filename string) (seeds []internal, in2outs []in2out) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	isFirst := true
	isParsingNum := false
	in2outIndex := -1
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		str := scanner.Text()
		if isFirst {
			s := strings.Split(str, ":")
			s = strings.Fields(s[1])
			for i := 0; i < len(s); i++ {
				n, err := strconv.ParseInt(s[i], 10, 64)
				if err != nil {
					panic(err)
				}
				seeds = append(seeds, internal(n))
			}
			isFirst = false
		} else {
			if str == "" {
				continue
			}
			if str[0]-'0' > 10 {
				isParsingNum = true
				in2outs = append(in2outs, in2out{
					make([][3]internal, 0),
				})
				in2outIndex++
				continue
			}
			if isParsingNum {
				s := strings.Fields(str)
				trio := [3]internal{}
				for i := 0; i < len(s); i++ {
					n, err := strconv.ParseInt(s[i], 10, 64)
					if err != nil {
						panic(err)
					}
					trio[i] = internal(n)
				}
				in2outs[in2outIndex].maps = append(in2outs[in2outIndex].maps, trio)
			}
		}
	}
	return
}

type in2out struct {
	maps [][3]internal
}

func (m *in2out) At(a internal) internal {
	for i := 0; i < len(m.maps); i++ {
		if a >= m.maps[i][1] && a < m.maps[i][1]+m.maps[i][2] {
			diff := a - m.maps[i][1]
			return m.maps[i][0] + diff
		}
	}
	return a
}

func walkThru(seed internal, maps []in2out) internal {
	n := seed
	for i := 0; i < len(maps); i++ {
		n = maps[i].At(n)
	}
	return n
}

func Solve(filename string) internal {
	seeds, other := parseInput(filename)
	fmt.Printf("%v\n", seeds)
	fmt.Printf("%v\n", other)
	results := []internal{}
	for i := 0; i < len(seeds); i++ {
		result := walkThru(seeds[i], other)
		results = append(results, result)
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i] < results[j]
	})
	return results[0]
}

func Solve2(filename string) internal {
	seeds, other := parseInput(filename)
	result := internal(-1)
	for i := 0; i < len(seeds); i += 2 {
		for j := internal(0); j < seeds[i+1]; j++ {
			tmp := walkThru(seeds[i]+j, other)
			if tmp < result {
				result = tmp
			} else if result == -1 {
				result = tmp
			}
		}
		fmt.Printf("seeds[i]=%d finished\n", seeds[i])
	}
	return result
}
