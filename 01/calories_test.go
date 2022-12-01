package main

import (
	"log"
	"os"
	"testing"
)

func TestCalculateTopCalories(t *testing.T) {
	t.Run("example case", func(t *testing.T) {
		file, err := os.Open("input1.txt")
		if err != nil {
			log.Fatal(err)
		}

		got := calculateTopCalories(file)
		want := 24000

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("exercise case", func(t *testing.T) {
		file, err := os.Open("input2.txt")
		if err != nil {
			log.Fatal(err)
		}

		got := calculateTopCalories(file)
		want := 67027

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestCalculateTop3Calories(t *testing.T) {
	t.Run("example case", func(t *testing.T) {
		file, err := os.Open("input1.txt")
		if err != nil {
			log.Fatal(err)
		}

		got := calculateTop3Calories(file)
		want := 41000

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("exercise case", func(t *testing.T) {
		file, err := os.Open("input2.txt")
		if err != nil {
			log.Fatal(err)
		}

		got := calculateTop3Calories(file)
		want := 197291

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestSumLastValues(t *testing.T) {
	cases := []struct {
		arr  []int
		num  int
		want int
	}{
		{[]int{1, 2, 3, 4}, 3, 9},
		{[]int{1, 2, 3, 4}, 1, 4},
		{[]int{1, 2, 3, 4}, 2, 7},
	}

	for _, tc := range cases {
		got := sumLastValues(tc.arr, tc.num)
		if got != tc.want {
			t.Errorf("got %v, want %v", got, tc.want)
		}
	}
}
