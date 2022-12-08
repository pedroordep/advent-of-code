package main

import "testing"

func TestCheckNonRepeatedCharIndex(t *testing.T) {
	t.Run("with num = 4", func(t *testing.T) {
		tt := []struct {
			got  []byte
			num  int
			want int
		}{
			{[]byte("bvwbjplbgvbhsrlpgdmjqwftvncz"), 4, 5},
			{[]byte("nppdvjthqldpwncqszvftbrmjlhg"), 4, 6},
			{[]byte("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"), 4, 10},
			{[]byte("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"), 4, 11},
		}

		for _, tc := range tt {
			if tc.want != checkNonRepeatedCharIndex(tc.got, tc.num) {
				t.Errorf("%v: got %d, want %d", string(tc.got), checkNonRepeatedCharIndex(tc.got, tc.num), tc.want)
			}
		}
	})

	t.Run("with num = 14", func(t *testing.T) {
		tt := []struct {
			got  []byte
			num  int
			want int
		}{
			{[]byte("mjqjpqmgbljsphdztnvjfqwrcgsmlb"), 14, 19},
			{[]byte("bvwbjplbgvbhsrlpgdmjqwftvncz"), 14, 23},
			{[]byte("nppdvjthqldpwncqszvftbrmjlhg"), 14, 23},
			{[]byte("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"), 14, 29},
			{[]byte("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"), 14, 26},
		}

		for _, tc := range tt {
			if tc.want != checkNonRepeatedCharIndex(tc.got, tc.num) {
				t.Errorf("%v: got %d, want %d", string(tc.got), checkNonRepeatedCharIndex(tc.got, tc.num), tc.want)
			}
		}
	})
}
