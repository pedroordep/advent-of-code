package main

import "testing"

func TestPart1(t *testing.T) {
	t.Run("should return 3", func(t *testing.T) {
		actual := Part1(`3-5
10-14
16-20
12-18

1
5
8
11
17
32`)
		expected := int64(3)

		if actual != expected {
			t.Errorf("got %d want %d", actual, expected)
		}
	})
}

func TestPart2(t *testing.T) {
	t.Run("should return 14", func(t *testing.T) {
		actual := Part2(`3-5
10-14
16-20
12-18

1
5
8
11
17
32`)
		expected := int64(14)

		if actual != expected {
			t.Errorf("got %d want %d", actual, expected)
		}
	})

	t.Run("should return 19", func(t *testing.T) {
		actual := Part2(`1-10
2-10
3-10
4-10
5-10
5-11
5-12
5-13
5-14
5-15
16-16
18-19
19-20

1
5
8
11
17
32`)
		expected := int64(19)

		if actual != expected {
			t.Errorf("got %d want %d", actual, expected)
		}
	})

	t.Run("should return 19", func(t *testing.T) {
		actual := Part2(`1-10
2-10
5-12
4-10
5-12
3-10
4-10
5-10
5-11
5-12
16-16
18-19
5-13
5-14
5-15
16-16
18-19
19-20

1
5
8
11
17
32`)
		expected := int64(19)

		if actual != expected {
			t.Errorf("got %d want %d", actual, expected)
		}
	})
}
