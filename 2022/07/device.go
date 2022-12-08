package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type Node struct {
	Parent   *Node
	Children []*Node
	Name     string
	Size     int
	IsDir    bool
}

func main() {
	file, _ := os.ReadFile("./input2.txt")

	split := strings.Split(string(file), "\n")

	tree := &Node{
		Parent:   nil,
		Children: make([]*Node, 0),
		Name:     "/",
		Size:     0,
		IsDir:    true,
	}

	setupNodeTree(tree, split)

	// fmt.Println(tree)

	updateDirSizes(tree)

	// fmt.Println(tree)

	fmt.Println("part 1:", sumOfDirSizes(tree, 100000))
	fmt.Println("part 2:", smallestDirToDelete(tree, 70000000-tree.Size, 30000000))
}

func sumOfDirSizes(n *Node, threshold int) int {
	sum := 0

	if !n.IsDir {
		return 0
	}
	if n.IsDir && n.Size < threshold {
		sum += n.Size
	}
	for _, child := range n.Children {
		sum += sumOfDirSizes(child, threshold)
	}

	return sum
}

func smallestDirToDelete(n *Node, unusedSize, updateSize int) int {
	dirNodeSizes := []int{}
	getNodeTreeDirSizes(n, &dirNodeSizes)

	sort.Ints(dirNodeSizes)

	for _, size := range dirNodeSizes {
		if size >= updateSize-unusedSize {
			return size
		}
	}

	return 0
}

func getNodeTreeDirSizes(n *Node, dirNodeSizes *[]int) {
	if n.IsDir {
		*dirNodeSizes = append(*dirNodeSizes, n.Size)
	}
	for _, child := range n.Children {
		getNodeTreeDirSizes(child, dirNodeSizes)
	}
}

func setupNodeTree(root *Node, commands []string) {
	curNode := root
	parsingDir := false
	for _, command := range commands[1:] {
		// fmt.Println(command)
		if parsingDir && command[0] != '$' {
			var name string
			if command[:3] == "dir" {
				fmt.Sscanf(command, "dir %s", &name)
				node := GetChildByName(curNode, name, true)
				if node != nil {
					break
				}
				curNode.Children = append(curNode.Children, &Node{
					Parent:   curNode,
					Children: make([]*Node, 0),
					Name:     name,
					IsDir:    true,
					Size:     0,
				})
			} else {
				var size int
				fmt.Sscanf(command, "%d %s", &size, &name)
				node := GetChildByName(curNode, name, false)
				if node != nil {
					break
				}
				curNode.Children = append(curNode.Children, &Node{
					Parent:   curNode,
					Children: nil,
					Name:     name,
					IsDir:    false,
					Size:     size,
				})
			}

			continue
		}
		parsingDir = false

		if command[:4] == "$ cd" {
			if strings.Contains(command, "/") {
				curNode = root
			} else if strings.Contains(command, "..") {
				curNode = curNode.Parent
			} else {
				var name string
				fmt.Sscanf(command, "$ cd %s", &name)
				for _, child := range curNode.Children {
					if child.IsDir && child.Name == name {
						curNode = child
						break
					}
				}
			}
		} else if command[:4] == "$ ls" {
			parsingDir = true
			continue
		}
	}
}

func GetChildByName(node *Node, name string, isDir bool) *Node {
	for _, child := range node.Children {
		if child.Name == name && child.IsDir == isDir {
			return child
		}
	}
	return nil
}

func updateDirSizes(n *Node) {
	sum := 0
	for _, child := range n.Children {
		if child.IsDir {
			updateDirSizes(child)
		}
		sum += child.Size
	}
	n.Size = sum
}

func (n *Node) String() string {
	buffer := bytes.Buffer{}
	printTree(&buffer, n, 0)
	return buffer.String()
}

func printTree(w io.Writer, n *Node, depth int) {
	prefix := strings.Repeat("  ", depth)
	fmt.Fprintf(w, "%s%s (%d)\n", prefix, n.Name, n.Size)
	for _, child := range n.Children {
		printTree(w, child, depth+1)
	}
}
