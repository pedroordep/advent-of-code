package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	t.Run("should return 50", func(t *testing.T) {
		actual := Part1(`7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`)
		expected := int(50)

		if actual != expected {
			t.Errorf("got %d want %d", actual, expected)
		}
	})
}

func TestPart2(t *testing.T) {
	t.Run("should return 24", func(t *testing.T) {
		actual := Part2(`7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`)
		expected := int(24)

		if actual != expected {
			t.Errorf("got %d want %d", actual, expected)
		}
	})

	t.Run("should return 56", func(t *testing.T) {
		actual := Part2(`2,6
4,6
4,4
6,4
6,6
8,6
8,2
10,2
10,4
12,4
12,6
14,6
14,3
16,3
16,8
18,8
18,5
20,5
20,10
15,10
15,9
5,9
5,10
2,10`)
		expected := int(56)

		if actual != expected {
			t.Errorf("got %d want %d", actual, expected)
		}
	})

	t.Run("should return 60", func(t *testing.T) {
		actual := Part2(`1,8
1,7
`)
		expected := int(56)

		if actual != expected {
			t.Errorf("got %d want %d", actual, expected)
		}
	})
}

// ..............
// .......#XXX#..
// .......XXXXX..
// ..OOOOOOOOXX..
// ..OOOOOOOOXX..
// ..OOOOOOOOXX..
// .........XXX..
// .........#X#..
// ..............

// ..............
// .......##.##..
// .......XX.XX..
// .......XX.XX..
// ..#XXXX##X#X..
// ..#XXXXXX#XX..
// .........XXX..
// .........#X#..
// ..............
