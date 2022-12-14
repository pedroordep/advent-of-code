package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Node struct {
	val    int
	parent *Node
	arr    Nodes
}
type Nodes []*Node

func (n Nodes) Len() int {
	return len(n)
}

func (n Nodes) Less(i, j int) bool {
	result := compareNodes(n[i], n[j])
	return result < 0
}

func (n Nodes) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n Nodes) IndexOf(node *Node) int {
	for i, v := range n {
		if node == v {
			return i
		}
	}
	return -1
}

func main() {
	file, _ := os.ReadFile("./input2.txt")
	split := strings.Split(string(file), "\n\n")

	fmt.Println("part 1:", part1(split))
	fmt.Println("part 1:", part2(split))
}

func part1(split []string) int {
	sum := 0
	for i, pair := range split {
		pairs := strings.Split(pair, "\n")

		first := parseNode(pairs[0])

		second := parseNode(pairs[1])

		// var first []any
		// json.Unmarshal([]byte(pairs[0]), &first)

		// var second []any
		// json.Unmarshal([]byte(pairs[1]), &second)

		val := compareNodes(first, second)
		// fmt.Println("Comparing", i+1)
		// fmt.Println(first, second, val)
		if val == -1 {
			sum += i + 1
		}
	}
	return sum
}

func part2(split []string) int {
	node1 := parseNode("[[2]]")
	node2 := parseNode("[[6]]")
	nodes := Nodes{node1, node2}
	for _, pair := range split {
		pairs := strings.Split(pair, "\n")

		nodes = append(nodes, parseNode(pairs[0]))
		nodes = append(nodes, parseNode(pairs[1]))
	}

	sort.Sort(nodes)

	return (nodes.IndexOf(node1) + 1) * (nodes.IndexOf(node2) + 1)
}

// -1 if first comes first
// 0 if equal
// 1 if second comes first
// func compareInputs(first, second any) int {
// 	fmt.Println(first, reflect.TypeOf(first), second, reflect.TypeOf(second))
// 	switch firstT := first.(type) {
// 	case float64:
// 		switch secondT := second.(type) {
// 		case float64:
// 			if firstT < secondT {
// 				return -1
// 			} else if firstT == second.(float64) {
// 				return 0
// 			} else {
// 				return 1
// 			}
// 		case []float64:
// 			return compareInputs([]float64{firstT}, secondT)
// 		}
// 	case []float64:
// 		switch secondT := second.(type) {
// 		case float64:
// 			return compareInputs(firstT, []float64{secondT})
// 		case []float64:
// 			for i := 0; i < len(firstT); i++ {
// 				if i == len(secondT) {
// 					return 1
// 				}
// 				val := compareInputs(firstT[i], secondT[i])
// 				if val != 0 {
// 					return val
// 				}
// 			}
// 			return 0
// 		}
// 	case []interface{}:
// 		switch secondT := second.(type) {
// 		case []interface{}:
// 			for i := 0; i < len(firstT); i++ {
// 				if i == len(secondT) {
// 					return 1
// 				}
// 				val := compareInputs(firstT[i].([]float64), secondT[i].([]float64))
// 				if val != 0 {
// 					return val
// 				}
// 			}
// 		case interface{}:
// 			return compareInputs(firstT, secondT)
// 		}
// 		return 0
// 	case interface{}:
// 		switch secondT := second.(type) {
// 		case []interface{}:
// 			return compareInputs([]interface{}{firstT}, secondT)
// 		case interface{}:
// 			return compareInputs(firstT, secondT)
// 		}
// 		return 0
// 	}
// 	return 0
// }

// -1 if first comes first
// 0 if equal
// 1 if second comes first
func compareNodes(first, second *Node) int {
	if first.val != -1 && second.val != -1 {
		if first.val < second.val {
			return -1
		} else if first.val == second.val {
			return 0
		} else {
			return 1
		}
	}
	if first.val == -1 && second.val == -1 {
		for i := 0; i < len(first.arr); i++ {
			if i == len(second.arr) {
				return 1
			}
			result := compareNodes(first.arr[i], second.arr[i])
			if result != 0 {
				return result
			}
		}
		if len(first.arr) < len(second.arr) {
			return -1
		}
	}
	if first.val != -1 {
		first.arr = append(first.arr, &Node{first.val, first, Nodes{}})
		first.val = -1
		return compareNodes(first, second)
	}
	if second.val != -1 {
		second.arr = append(second.arr, &Node{second.val, second, Nodes{}})
		second.val = -1
		return compareNodes(first, second)
	}
	return 0
}

func parseNode(input string) *Node {
	root := &Node{-1, nil, Nodes{}}
	cur := root
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '[':
			new := &Node{-1, cur, Nodes{}}
			cur.arr = append(cur.arr, new)
			cur = new
		case ']':
			cur = cur.parent
		case ',':
			// skip?
		default:
			j := i
			str := ""
			for input[j] != ',' && input[j] != ']' {
				str += string(input[j])
				j++
			}
			val, _ := strconv.Atoi(str)
			cur.arr = append(cur.arr, &Node{val, cur, nil})
		}
	}

	return root
}

func (n *Node) String() string {
	return printRecursive(n)
}

func printRecursive(n *Node) string {
	if n.val != -1 {
		return fmt.Sprint(n.val)
	}
	if len(n.arr) > 0 {
		str := "["
		for i, v := range n.arr {
			str += printRecursive(v)
			if i < len(n.arr)-1 {
				str += ","
			}
		}
		return str + "]"
	}
	return ""
}
