package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Valve struct {
	id          string
	flow        int
	isOpen      bool
	connections []*Valve
	directCons  []*Connection
}

type Connection struct {
	to      *Valve
	minutes int
}

type ValveSet map[*Valve]struct{}
type ValveMap map[string]*Valve

var bestGlobal int = 0
var maxScorePerTurn int = 0
var valves ValveMap
var valvesOpened ValveSet
var closestDistances map[*Valve]int

func main() {
	file, _ := os.ReadFile("./input2.txt")
	split := strings.Split(string(file), "\n")
	root := parseGraph(split, "AA")

	// printValves(valves)
	fmt.Println(valves)
	closestDistances = Dijkstra(root, true)
	// fmt.Println(closestDistances)

	fmt.Println("part 1:", bestPath(root, nil, 30, 0, []string{}))
}

func bestPath(cur *Valve, con *Connection, minutes, score int, path []string) int {
	// fmt.Println(path, len(path), minutes, bestGlobal)
	bestLocal := 0

	if con != nil {
		if minutes-con.minutes < 0 {
			score += sumScores() * minutes
			minutes = 0
		} else {
			score += sumScores() * con.minutes
			minutes -= con.minutes
		}
	} else {
		score += sumScores()
		minutes -= 1
	}

	// base cases

	if minutes <= 0 {
		return score
	}
	if len(valvesOpened) == len(valves) {
		return score + sumScores()*minutes
	}
	if score+maxScoreFromValve(cur, minutes) < bestGlobal {
		return score
	}

	// navigation
	_, exists := valvesOpened[cur]
	if !exists {
		cur.isOpen = true
		valvesOpened[cur] = struct{}{}
		val := bestPath(cur, nil, minutes, score, append(path, "OPEN"))
		if val > bestLocal {
			bestLocal = val
		}
		if val > bestGlobal {
			bestGlobal = val
			fmt.Println(path, minutes, sumScores(), bestGlobal, maxScorePerTurn*30)
		}
		cur.isOpen = false
		delete(valvesOpened, cur)
	}

	for _, con := range cur.directCons {
		val := bestPath(con.to, con, minutes, score, append(path, con.to.id))
		if val > bestLocal {
			bestLocal = val
		}
		if val > bestGlobal {
			bestGlobal = val
			fmt.Println(path, minutes, sumScores(), bestGlobal, maxScorePerTurn*30)
		}
	}

	return bestLocal
}

func sumScores() int {
	sum := 0
	for _, valve := range valves {
		if valve.isOpen {
			sum += valve.flow
		}
	}
	return sum
}

func maxScoreFromValve(valve *Valve, minutes int) int {
	maxScore := 0
	distances := Dijkstra(valve, true)
	for v, d := range distances {
		if v.isOpen {
			maxScore += v.flow * (minutes)
		} else if minutes-d > 0 {
			maxScore += v.flow * (minutes - d)
		}
	}
	return maxScore
}

func parseGraph(split []string, rootId string) *Valve {
	// create valves
	valves = ValveMap{}
	valvesOpened = ValveSet{}
	for _, line := range split {
		id, flow := "", 0
		fmt.Sscanf(line, "Valve %s has flow rate=%d;", &id, &flow)

		newValve := &Valve{id, flow, false, []*Valve{}, []*Connection{}}
		if flow == 0 {
			newValve.isOpen = true
			valvesOpened[newValve] = struct{}{}
		}
		maxScorePerTurn += flow
		valves[id] = newValve
	}

	// build graph and set root
	var root *Valve
	for _, line := range split {
		valveInfo := strings.Split(line, "; ")
		id, flow := "", 0
		fmt.Sscanf(valveInfo[0], "Valve %s has flow rate=%d;", &id, &flow)

		if id == rootId {
			root = valves[id]
		}

		valveInfo[1] = strings.ReplaceAll(valveInfo[1], "tunnels lead to valves ", "")
		valveInfo[1] = strings.ReplaceAll(valveInfo[1], "tunnel leads to valve ", "")
		valveIds := strings.Split(valveInfo[1], ", ")
		for _, connectionId := range valveIds {
			valves[id].connections = append(valves[id].connections, valves[connectionId])
		}
	}

	// populate direct connections on valves
	for _, valve := range valves {
		distances := Dijkstra(valve, false)

		for v, d := range distances {
			if v != valve && (v.flow > 0 || v == root) {
				valve.directCons = append(valve.directCons, &Connection{v, d})
			}
		}
	}

	// remove valves with flow 0 and not the root
	for id, valve := range valves {
		if valve.flow == 0 && valve != root {
			delete(valves, id)
			delete(valvesOpened, valve)
		}
	}

	return root
}

func Dijkstra(start *Valve, withWeights bool) map[*Valve]int {
	dist := map[*Valve]int{}
	sptSet := map[*Valve]bool{}

	for _, valve := range valves {
		dist[valve] = math.MaxInt32
		sptSet[valve] = false
	}

	dist[start] = 0

	for count := 0; count < len(valves)-1; count++ {
		u := minDistance(dist, sptSet)

		sptSet[u] = true

		for v, _ := range dist {
			if !withWeights {
				if !sptSet[v] && u.connectsTo(v) && dist[u] != math.MaxInt32 && dist[u]+1 < dist[v] {
					dist[v] = dist[u] + 1
				}
			} else {
				index := u.directConnectToIndex(v)
				if !sptSet[v] && index != -1 && dist[u] != math.MaxInt32 && dist[u]+u.directCons[index].minutes < dist[v] {
					dist[v] = dist[u] + u.directCons[index].minutes
				}
			}
		}
	}

	return dist
}

func (v *Valve) connectsTo(t *Valve) bool {
	for _, connection := range v.connections {
		if connection == t {
			return true
		}
	}
	return false
}

func (v *Valve) directConnectToIndex(t *Valve) int {
	for i, connection := range v.directCons {
		if connection.to == t {
			return i
		}
	}
	return -1
}

func minDistance(dist map[*Valve]int, sptSet map[*Valve]bool) *Valve {
	min := math.MaxInt32
	var minValve *Valve

	for valve, distance := range dist {
		if !sptSet[valve] && distance <= min {
			min = dist[valve]
			minValve = valve
		}
	}
	return minValve
}

func printValves(valves ValveMap) {
	for _, valve := range valves {
		conStr := []string{}
		for _, vc := range valve.connections {
			conStr = append(conStr, vc.id)
		}
		fmt.Printf("Valve %s has flow rate=%d (%v); tunnels lead to valves %s\n", valve.id, valve.flow, valve.isOpen, strings.Join(conStr, ", "))
	}
}

func (v *Valve) String() string {
	return fmt.Sprintf("{id: %s, flow: %d, isOpen: %t, %v}", v.id, v.flow, v.isOpen, v.directCons)
}

func (c *Connection) String() string {
	return fmt.Sprintf("{to: %s, dist: %d}", c.to.id, c.minutes)
}
