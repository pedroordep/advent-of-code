package main

import (
	"fmt"
	"math/big"
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
	items     []*big.Int
	op        int
	opTerm    *big.Int // 0 means the old param
	testDiv   int
	testTrue  int
	testFalse int
}

var lowestDivValue int64 = 1

func main() {
	file, _ := os.ReadFile("./input2.txt")
	split := strings.Split(string(file), "\n\n")

	monkeys := parseMonkeys(split)
	monkeyInspections := calculateInspections(monkeys, 20, true)

	monkeys = parseMonkeys(split)
	monkeyInspections2 := calculateInspections(monkeys, 10000, false)

	fmt.Println("part 1", monkeyInspections)
	fmt.Println("part 2", monkeyInspections2)

	// fmt.Println(big.NewInt(lowestDivValue), big.NewInt(0).Mul(big.NewInt(lowestDivValue), big.NewInt(lowestDivValue)))
}

func calculateInspections(monkeys []*Monkey, rounds int, worryLvlChange bool) int {
	monkeyInspections := make([]int, len(monkeys))

	for round := 0; round < rounds; round++ {
		for i := 0; i < len(monkeys); i++ {
			if len(monkeys[i].items) == 0 {
				continue
			}

			for _, item := range monkeys[i].items {
				// inspect
				monkeyInspections[i]++
				// fmt.Printf("[%d]Monkey inspects an item with a worry level of %d.\n", i, item)

				// worry level calc
				worryLvl := big.NewInt(0)
				term := big.NewInt(0)
				if monkeys[i].opTerm.Cmp(big.NewInt(0)) == 0 {
					term.Set(item)
				} else {
					term.Set(monkeys[i].opTerm)
				}
				if monkeys[i].op == opMult {
					worryLvl.Mul(item, term)
					// fmt.Printf("[%d]  Worry level is multiplied by %d to %d.\n", i, term, worryLvl)
				} else {
					worryLvl.Add(item, term)
					// fmt.Printf("[%d]  Worry level is added by %d to %d.\n", i, term, worryLvl)
				}

				// bored
				if worryLvlChange {
					worryLvl.Div(worryLvl, big.NewInt(3))
					// fmt.Printf("[%d]  Monkey gets bored with item. Worry level is divided by 3 to %d.\n", i, worryLvl)
				}

				// reset to %lowestDivValue
				worryLvl.Mod(worryLvl, big.NewInt(lowestDivValue))

				// throw item
				if big.NewInt(0).Mod(worryLvl, big.NewInt(int64(monkeys[i].testDiv))).Cmp(big.NewInt(0)) == 0 {
					monkeys[monkeys[i].testTrue].items = append(monkeys[monkeys[i].testTrue].items, worryLvl)
					// fmt.Printf("[%d]    Item with worry level %d is thrown to monkey %d.\n", i, worryLvl, monkeys[i].testTrue)
				} else {
					monkeys[monkeys[i].testFalse].items = append(monkeys[monkeys[i].testFalse].items, worryLvl)
					// fmt.Printf("[%d]    Item with worry level %d is thrown to monkey %d.\n", i, worryLvl, monkeys[i].testFalse)
				}
			}
			monkeys[i].items = nil
		}
	}

	sort.Ints(monkeyInspections)
	// fmt.Println(monkeyInspections)

	return monkeyInspections[len(monkeyInspections)-1] * monkeyInspections[len(monkeyInspections)-2]
}

func parseMonkeys(input []string) []*Monkey {
	monkeys := []*Monkey{}
	for i := 0; i < len(input); i++ {
		lines := strings.Split(input[i], "\n")
		monkey := &Monkey{
			id:    i,
			items: make([]*big.Int, 0),
		}

		// starting items
		itemList := strings.Split(strings.TrimPrefix(lines[1], "  Starting items: "), ", ")
		for _, item := range itemList {
			val, _ := strconv.Atoi(item)
			monkey.items = append(monkey.items, big.NewInt(int64(val)))
		}

		// operation
		opList := strings.TrimPrefix(lines[2], "  Operation: new = old ")
		if opList[0] == '*' {
			monkey.op = opMult
		} else {
			monkey.op = opAdd
		}
		var opTerm int64
		fmt.Sscanf(opList[2:], "%d", &opTerm)
		monkey.opTerm = big.NewInt(opTerm)

		// test
		fmt.Sscanf(lines[3], "  Test: divisible by %d", &monkey.testDiv)
		fmt.Sscanf(lines[4], "    If true: throw to monkey %d", &monkey.testTrue)
		fmt.Sscanf(lines[5], "    If false: throw to monkey %d", &monkey.testFalse)
		lowestDivValue *= int64(monkey.testDiv)

		monkeys = append(monkeys, monkey)
	}

	return monkeys
}

func (m *Monkey) String() string {
	return fmt.Sprintf("[%d] items: %v, op: %d %d, test: %%%d ? %d : %d", m.id, m.items, m.op, m.opTerm, m.testDiv, m.testTrue, m.testFalse)
}
