package main

import (
	"container/heap"
	"errors"
	"fmt"
	"image"
	"os"
	"strings"
)

// type Array []image.Point

func main() {
	file, _ := os.ReadFile("./input2.txt")
	maze := strings.Split(string(file), "\n")

	_, start, end, starts := parseInput(maze)

	path, _ := ShortestPath(start, end)
	fmt.Println("part 1:", len(path)-1)

	fmt.Println("part 2:", minOfMultiplePaths(starts, end))
}

func minOfMultiplePaths(starts []Node, end Node) int {
	min := 999999
	for _, start := range starts {
		path, err := ShortestPath(start, end)
		if err != nil {
			continue
		}
		if len(path) < min {
			min = len(path)
		}
	}
	return min - 1
}

func parseInput(maze []string) (graph []Node, start, end Node, starts []Node) {
	graph = []Node{}
	boundaries := image.Rect(0, 0, len(maze[0]), len(maze))
	moves := []image.Point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	vertexes := map[image.Point]*Vertex{}
	starts = []Node{}

	for y := 0; y < len(maze); y++ {
		for x := 0; x < len(maze[y]); x++ {
			point := image.Point{x, y}
			v := &Vertex{x, y, []*Link{}}
			vertexes[point] = v

			if maze[y][x] == 'S' {
				start = v
				starts = append(starts, v)
			} else if maze[y][x] == 'E' {
				end = v
			}
			if maze[y][x] == 'a' {
				starts = append(starts, v)
			}
		}
	}
	for y := 0; y < len(maze); y++ {
		for x := 0; x < len(maze[y]); x++ {
			point := image.Point{x, y}
			graph = append(graph, vertexes[point])

			for _, move := range moves {
				toPoint := point.Add(move)

				if toPoint.In(boundaries) {
					fromElevation := maze[point.Y][point.X]
					toElevation := maze[toPoint.Y][toPoint.X]
					if maze[point.Y][point.X] == 'S' {
						fromElevation = 'a'
					}
					if maze[toPoint.Y][toPoint.X] == 'E' {
						toElevation = 'z'
					}

					if toElevation <= fromElevation+1 {
						l := &Link{vertexes[toPoint]}
						vertexes[point].Links = append(vertexes[point].Links, l)
					}
				}
			}
		}
	}

	return graph, start, end, starts
}

// Node is an interface for your own implementation of a vertex in a graph.
type Node interface {
	Edges() []Edge
}

type Vertex struct {
	X, Y  int
	Links []*Link
}

func (v *Vertex) Edges() []Edge {
	edges := []Edge{}
	for _, link := range v.Links {
		edges = append(edges, link)
	}
	return edges
}

// Edge is an interface for your own implementation of an edge between two vertices in a graph.
type Edge interface {
	Destination() Node
	Weight() float64
}

type Link struct {
	To *Vertex
}

func (v *Link) Destination() Node {
	return v.To
}

func (v *Link) Weight() float64 {
	return 1
}

func ShortestPath(start, end Node) ([]Node, error) {
	visited := make(map[Node]struct{})
	dists := make(map[Node]float64)
	prev := make(map[Node]Node)

	dists[start] = 0
	queue := &queue{&queueItem{value: start, weight: 0, index: 0}}
	heap.Init(queue)

	for queue.Len() > 0 {
		// Done.
		if _, ok := visited[end]; ok {
			break
		}

		item := heap.Pop(queue).(*queueItem)
		n := item.value
		for _, edge := range n.Edges() {
			dest := edge.Destination()
			dist := dists[n] + edge.Weight()
			if tentativeDist, ok := dists[dest]; !ok || dist < tentativeDist {
				dists[dest] = dist
				prev[dest] = n
				heap.Push(queue, &queueItem{value: dest, weight: dist})
			}
		}
		visited[n] = struct{}{}
	}

	if _, ok := visited[end]; !ok {
		return nil, errors.New("no shortest path exists")
	}

	path := []Node{end}
	for next := prev[end]; next != nil; next = prev[next] {
		path = append(path, next)
	}

	// Reverse path.
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path, nil
}

type queueItem struct {
	value  Node
	weight float64
	index  int
}

type queue []*queueItem

func (q queue) Len() int {
	return len(q)
}

func (q queue) Less(i, j int) bool {
	return q[i].weight < q[j].weight
}

func (q queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

func (q *queue) Push(x interface{}) {
	n := len(*q)
	item := x.(*queueItem)
	item.index = n
	*q = append(*q, item)
}

func (q *queue) Pop() interface{} {
	old := *q
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*q = old[0 : n-1]
	return item
}
