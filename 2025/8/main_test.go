package main

import "testing"

func TestPart1(t *testing.T) {
	t.Run("should return 40", func(t *testing.T) {
		actual := Part1(`162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`, 10)
		expected := int(40)

		if actual != expected {
			t.Errorf("got %d want %d", actual, expected)
		}
	})

	t.Run("should return 25272", func(t *testing.T) {
		actual := Part2(`162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`, -1)
		expected := int(25272)

		if actual != expected {
			t.Errorf("got %d want %d", actual, expected)
		}
	})
}
