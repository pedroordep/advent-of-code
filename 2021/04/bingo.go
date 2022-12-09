package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BingoCell struct {
	val    int
	marked bool
}
type BingoCard [][]*BingoCell

func main() {
	file, _ := os.ReadFile("./input2.txt")
	split := strings.Split(string(file), "\n\n")

	numbers := strings.Split(split[0], ",")
	bingoCards := parseBingoCards(split[1:])

	card, number := playBingo(numbers, bingoCards)

	fmt.Println("part 1:", calculatePoints(card)*number)

	card, number = playBingoToLose(numbers, bingoCards)
	fmt.Println(card, number)

	fmt.Println("part 2:", calculatePoints(card)*number)
}

func calculatePoints(card *BingoCard) int {
	sum := 0
	for i := 0; i < len(*card); i++ {
		for j := 0; j < len((*card)[i]); j++ {
			if !(*card)[i][j].marked {
				sum += (*card)[i][j].val
			}
		}
	}
	return sum
}

func playBingo(numbersStr []string, cards []*BingoCard) (*BingoCard, int) {
	for _, numberStr := range numbersStr {
		number, _ := strconv.Atoi(numberStr)
		for _, card := range cards {
			if markNumber(card, number) {
				return card, number
			}
		}
	}

	return &BingoCard{}, 0
}

func playBingoToLose(numbersStr []string, cards []*BingoCard) (*BingoCard, int) {
	for _, numberStr := range numbersStr {
		number, _ := strconv.Atoi(numberStr)
		for i := 0; i < len(cards); i++ {
			if markNumber(cards[i], number) {
				if len(cards) == 1 {
					return cards[i], number
				} else {
					// fmt.Println("removing card\n", cards[i])
					removeCardFromCards(&cards, cards[i])
					i -= 1
				}
			}
		}
	}

	return &BingoCard{}, 0
}

func markNumber(card *BingoCard, number int) (bingo bool) {
	bingo = false
	for i := 0; i < len(*card); i++ {
		for j := 0; j < len((*card)[i]); j++ {
			if (*card)[i][j].val == number {
				(*card)[i][j].marked = true
				return checkBingo(card, i, j)
			}
		}
	}
	return bingo
}

func checkBingo(card *BingoCard, x, y int) bool {
	line, column := true, true
	for i := 0; i < len(*card); i++ {
		if !(*card)[i][y].marked {
			column = false
			break
		}
	}
	for j := 0; j < len((*card)[0]); j++ {
		if !(*card)[x][j].marked {
			line = false
			break
		}
	}
	return line || column
}

func parseBingoCards(inputs []string) []*BingoCard {
	cards := []*BingoCard{}
	for _, input := range inputs {
		card := &BingoCard{}
		lines := strings.Split(input, "\n")
		for _, line := range lines {
			row := []*BingoCell{
				{0, false},
				{0, false},
				{0, false},
				{0, false},
				{0, false},
			}
			fmt.Sscanf(line, "%d %d %d %d %d", &row[0].val, &row[1].val, &row[2].val, &row[3].val, &row[4].val)
			*card = append(*card, row)
		}
		cards = append(cards, card)
	}

	return cards
}

func removeCardFromCards(cards *[]*BingoCard, card *BingoCard) {
	for i := 0; i < len(*cards); i++ {
		if card == (*cards)[i] {
			*cards = append((*cards)[:i], (*cards)[i+1:]...)
			return
		}
	}
}

func (bc *BingoCard) String() string {
	str := ""
	for i := 0; i < len(*bc); i++ {
		for j := 0; j < len((*bc)[i]); j++ {
			marked := ""
			if (*bc)[i][j].marked {
				marked = "X"
			} else {
				marked = "_"
			}
			str += fmt.Sprintf("%2d%s ", (*bc)[i][j].val, marked)
		}
		str += "\n"
	}
	return str
}
