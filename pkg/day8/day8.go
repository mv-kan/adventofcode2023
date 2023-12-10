package day8

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type node struct {
	l    *node
	elem string
	r    *node
}

func readFile(filename string) string {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(file)
}

func parseInput(filename string) (beginEnd []string, cmd string, nodes map[string][]string) {
	nodes = map[string][]string{}
	filecontent := readFile(filename)
	lines := strings.Split(filecontent, "\n")
	isFirst := true
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for i := 0; i < len(lines); i++ {
		str := lines[i]
		if isFirst {
			cmd = str
			isFirst = false
		} else {
			if str == "" {
				continue
			}
			s := strings.Split(str, "=")
			key := strings.Trim(s[0], " ")
			n := strings.Trim(s[1], " ()")
			node := strings.Split(n, ",")
			nodes[key] = []string{node[0], strings.Trim(node[1], " ")}
			if len(beginEnd) == 0 {
				beginEnd = append(beginEnd, key)
			}
			if len(lines) == i+1 {
				beginEnd = append(beginEnd, key)
			}
		}
	}
	return
}

func fillGraph(currNode *node, beginEnd []string, maps map[string][]string, elem2node map[string]*node) {
	// if curr node is the end then return
	if currNode.r != nil {
		return
	}
	s := maps[currNode.elem]
	left := s[0]
	right := s[1]

	currNode.l = elem2node[left]
	currNode.r = elem2node[right]

	fillGraph(currNode.l, beginEnd, maps, elem2node)
	fillGraph(currNode.r, beginEnd, maps, elem2node)
}

func createGraph(beginEnd []string, maps map[string][]string) *node {
	elem2node := map[string]*node{}
	for k := range maps {
		elem2node[k] = &node{
			elem: k,
		}
	}
	firstNode := elem2node[beginEnd[0]]
	fillGraph(firstNode, beginEnd, maps, elem2node)
	return firstNode
}

func Solve(filename string) int {
	beginEnd, cmd, m := parseInput(filename)
	beginEnd[0] = "AAA"
	beginEnd[1] = "ZZZ"
	graph := createGraph(beginEnd, m)
	result := 0
	currNode := graph
	for i := 0; ; i++ {
		if i >= len(cmd) {
			i = 0
		}
		fmt.Printf("%s", string(cmd[i]))
		// fmt.Printf("%s -> (%s) ", currNode.elem, string(cmd[i]))
		if currNode.elem == beginEnd[1] {
			fmt.Println(i)
			break
		}
		if cmd[i] == 'L' {
			currNode = currNode.l
		} else if cmd[i] == 'R' {
			currNode = currNode.r
		}
		result++
	}
	return result
}
