package main

import (
	"fmt"
	"os"
	"strings"
)

type Shape struct {
	Points       int
	WinsAgainst  *Shape
	LosesAgainst *Shape
}

var rock *Shape
var paper *Shape
var scissors *Shape

func main() {
	file, err := os.ReadFile("./input2.txt")
	if err != nil {
		panic(err)
	}

	split := strings.Split(string(file), "\n")

	rock = &Shape{Points: 1}
	paper = &Shape{Points: 2}
	scissors = &Shape{Points: 3}
	rock.WinsAgainst = scissors
	rock.LosesAgainst = paper
	paper.WinsAgainst = rock
	paper.LosesAgainst = scissors
	scissors.WinsAgainst = paper
	scissors.LosesAgainst = rock

	result1 := firstGame(split)
	result2 := secondGame(split)

	fmt.Println("firstGame:", result1)
	fmt.Println("secondGame:", result2)
}

func firstGame(split []string) int {
	letters := map[byte]*Shape{
		'A': rock,
		'X': rock,
		'B': paper,
		'Y': paper,
		'C': scissors,
		'Z': scissors,
	}

	sum := 0

	for _, game := range split {
		// points from what you played
		sum += letters[game[2]].Points

		// points from the result of the match
		if letters[game[2]] == letters[game[0]] {
			sum += 3
		} else if letters[game[2]].WinsAgainst == letters[game[0]] {
			sum += 6
		}
	}

	return sum
}

func secondGame(split []string) int {
	var lose byte = 'X'
	var draw byte = 'Y'
	var win byte = 'Z'

	opponent := map[byte]*Shape{
		'A': rock,
		'B': paper,
		'C': scissors,
	}

	outcomes := map[byte]int{
		lose: 0,
		draw: 3,
		win:  6,
	}

	sum := 0

	for _, game := range split {
		fmt.Println(game)
		// points from the outcome of the match
		sum += outcomes[game[2]]

		// points from what you'll play
		if game[2] == lose {
			sum += opponent[game[0]].WinsAgainst.Points
		} else if game[2] == win {
			sum += opponent[game[0]].LosesAgainst.Points
		} else {
			sum += opponent[game[0]].Points
		}
		fmt.Println(sum)
	}

	return sum
}
