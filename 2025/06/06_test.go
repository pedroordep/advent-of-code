package main

import "testing"

func TestPart1(t *testing.T) {
	t.Run("should return 4277556", func(t *testing.T) {
		actual := Part1(`123 328  51 64
 45 64  387 23
  6 98  215 314
*   +   *   +  `)
		expected := int64(4277556)

		if actual != expected {
			t.Errorf("got %d want %d", actual, expected)
		}
	})
}

func TestPart2(t *testing.T) {
	t.Run("should return 3263827", func(t *testing.T) {
		actual := Part2("123 328  51 64 \n 45 64  387 23 \n  6 98  215 314\n*   +   *   +  ")

		expected := int64(3263827)

		if actual != expected {
			t.Errorf("got %d want %d", actual, expected)
		}
	})
}
