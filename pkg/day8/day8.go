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

func fillGraph(currNode *node, maps map[string][]string, elem2node map[string]*node) {
	// if curr node is the end then return
	if currNode.r != nil {
		return
	}
	s := maps[currNode.elem]
	left := s[0]
	right := s[1]

	currNode.l = elem2node[left]
	currNode.r = elem2node[right]

	fillGraph(currNode.l, maps, elem2node)
	fillGraph(currNode.r, maps, elem2node)
}

func createGraph(beginEnd []string, maps map[string][]string) *node {
	elem2node := map[string]*node{}
	for k := range maps {
		elem2node[k] = &node{
			elem: k,
		}
	}
	firstNode := elem2node[beginEnd[0]]
	fillGraph(firstNode, maps, elem2node)
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

func parseInput2(filename string) (begins []string, cmd string, nodes map[string][]string) {
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
			if key[2] == 'A' {
				begins = append(begins, key)
			}
		}
	}
	return
}

func createGraph2(begins []string, maps map[string][]string) []*node {
	elem2node := map[string]*node{}
	for k := range maps {
		elem2node[k] = &node{
			elem: k,
		}
	}
	nodes := []*node{}
	for _, v := range begins {
		n := elem2node[v]
		fillGraph(n, maps, elem2node)
		nodes = append(nodes, n)
	}
	return nodes
}

func stepThru(nodes []*node, cmd byte) {
	for i, n := range nodes {
		if cmd == 'L' {
			nodes[i] = n.l
		} else if cmd == 'R' {
			nodes[i] = n.r
		}
	}
}

func checkForLastZ(nodes []*node) (bool, int) {
	for i, n := range nodes {
		if n.elem[2] == 'Z' {
			return true, i
		}
	}
	return false, 0
}

// to solve it you need to calculate LCD for steps values
func Solve2(filename string) int {
	begins, cmd, m := parseInput2(filename)
	nodes := createGraph2(begins, m)
	ogNodes := []node{}
	for i := 0; i < len(nodes); i++ {
		ogNodes = append(ogNodes, *nodes[i])
	}
	result := 0
	for i := 0; ; i++ {
		if i >= len(cmd) {
			i = 0
		}
		// fmt.Printf("%s", string(cmd[i]))
		stepThru(nodes, cmd[i])
		result++
		z, index := checkForLastZ(nodes)
		if z {

			fmt.Printf("og=%s, steps=%d, elem=%s\n", ogNodes[index].elem, result, nodes[index].elem)
			ogNodes[index].elem = ""
			fmt.Println()
			// gotToZ = true
		}
		if result > 100000 {
			return 0
		}
	}
}
