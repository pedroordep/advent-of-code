package main

import "testing"

func TestPart1(t *testing.T) {
	t.Run("should return 13", func(t *testing.T) {
		actual := Part1(`..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`)
		expected := 13

		if actual != expected {
			t.Errorf("got %d want %d", actual, expected)
		}
	})

	t.Run("should return 4", func(t *testing.T) {
		actual := getAdjacentRolls(0, 7, []string{"..@@.@@@@.", "@@@.@.@.@@"})

		expected := 4

		if actual != expected {
			t.Errorf("got %d want %d", actual, expected)
		}
	})
}

func TestPart2(t *testing.T) {
	t.Run("should return 43", func(t *testing.T) {
		actual := Part2(`..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`)
		expected := int(43)

		if actual != expected {
			t.Errorf("got %d want %d", actual, expected)
		}
	})
}
