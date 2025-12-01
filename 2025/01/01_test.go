package main

import "testing"

func TestPart1(t *testing.T) {
	t.Run("should return 3", func(t *testing.T) {
		actual := Part1(
			`L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`)
		expected := 3

		if actual != expected {
			t.Errorf("got %d want %d", actual, expected)
		}
	})
}

func TestPart2(t *testing.T) {
	t.Run("should return 6", func(t *testing.T) {
		actual := Part2(
			`L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`)
		expected := 6

		if actual != expected {
			t.Errorf("got %d want %d", actual, expected)
		}
	})

	var tableTests = []struct {
		input    string
		expected int
	}{
		{"L50\nL500", 6},
		{"L50\nL1\nR2\nL1\nR2", 3},
		{"R1000", 10},
		{"R50\nL1\nL99\nL1\nL100", 3},
		{"R49\nL98\nR98", 0},
		{"R49\nL199\nR199", 3},
	}

	for _, tt := range tableTests {
		t.Run(tt.input, func(t *testing.T) {
			actual := Part2(tt.input)
			if actual != tt.expected {
				t.Errorf("got %v, want %v", actual, tt.expected)
			}
		})
	}
}
