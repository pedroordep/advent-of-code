package main

import "testing"

func TestPart1(t *testing.T) {
	t.Run("should return 357", func(t *testing.T) {
		actual := Part1(`987654321111111
811111111111119
234234234234278
818181911112111`)
		expected := 357

		if actual != expected {
			t.Errorf("got %d want %d", actual, expected)
		}
	})
}

func TestPart2(t *testing.T) {
	t.Run("should return 4174379265", func(t *testing.T) {
		actual := Part2(`987654321111111
811111111111119
234234234234278
818181911112111`)
		expected := int64(3121910778619)

		if actual != expected {
			t.Errorf("got %d want %d", actual, expected)
		}
	})
}
