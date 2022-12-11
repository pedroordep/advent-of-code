package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	opAdd = iota
	opMult
)

type Monkey struct {
	id        int
	items     []int
	op        int
	opTerm    int // 0 means the old param
	testDiv   int
	testTrue  int
	testFalse int
}

func main() {
	file, _ := os.ReadFile("./input1.txt")
	split := strings.Split(string(file), "\n\n")

	monkeys := parseMonkeys(split)
	// fmt.Println(monkeys)
	monkeyInspections := calculateInspections(monkeys, 20, true)
	monkeyInspections2 := calculateInspections(monkeys, 10000, false)

	fmt.Println("part 1", monkeyInspections)
	fmt.Println("part 2", monkeyInspections2)
}

func calculateInspections(monkeys []*Monkey, rounds int, worryLvlChange bool) int {
	monkeyInspections := make([]int, len(monkeys))

	for round := 0; round < rounds; round++ {
		for i := 0; i < len(monkeys); i++ {
			if len(monkeys[i].items) == 0 {
				continue
			}

			itemsCp := make([]int, len(monkeys[i].items))
			copy(itemsCp, monkeys[i].items)

			for _, item := range itemsCp {
				// inspect
				monkeyInspections[i]++

				// worry level calc
				worryLvl, term := 0, 0
				if monkeys[i].opTerm == 0 {
					term = item
				} else {
					term = monkeys[i].opTerm
				}
				if monkeys[i].op == opMult {
					worryLvl = item * term
				} else {
					worryLvl = item + term
				}

				// bored
				if worryLvlChange {
					worryLvl = worryLvl / 3
				}

				// throw item
				if worryLvl%monkeys[i].testDiv == 0 {
					monkeys[monkeys[i].testTrue].items = append(monkeys[monkeys[i].testTrue].items, worryLvl)
				} else {
					monkeys[monkeys[i].testFalse].items = append(monkeys[monkeys[i].testFalse].items, worryLvl)
				}
				monkeys[i].items = monkeys[i].items[1:]
			}
		}
	}

	sort.Ints(monkeyInspections)
	fmt.Println(monkeyInspections)

	return monkeyInspections[len(monkeyInspections)-1] * monkeyInspections[len(monkeyInspections)-2]
}

func parseMonkeys(input []string) []*Monkey {
	monkeys := []*Monkey{}
	for i := 0; i < len(input); i++ {
		lines := strings.Split(input[i], "\n")
		monkey := &Monkey{
			id:    i,
			items: make([]int, 0),
		}

		// starting items
		itemList := strings.Split(strings.TrimPrefix(lines[1], "  Starting items: "), ", ")
		for _, item := range itemList {
			val, _ := strconv.Atoi(item)
			monkey.items = append(monkey.items, val)
		}

		// operation
		opList := strings.TrimPrefix(lines[2], "  Operation: new = old ")
		if opList[0] == '*' {
			monkey.op = opMult
		} else {
			monkey.op = opAdd
		}
		fmt.Sscanf(opList[1:], " %d", &monkey.opTerm)

		// test
		fmt.Sscanf(lines[3], "  Test: divisible by %d", &monkey.testDiv)
		fmt.Sscanf(lines[4], "    If true: throw to monkey %d", &monkey.testTrue)
		fmt.Sscanf(lines[5], "    If false: throw to monkey %d", &monkey.testFalse)

		monkeys = append(monkeys, monkey)
	}

	return monkeys
}

func (m *Monkey) String() string {
	return fmt.Sprintf("[%d] items: %v, op: %d %d, test: %%%d ? %d : %d", m.id, m.items, m.op, m.opTerm, m.testDiv, m.testTrue, m.testFalse)
}
