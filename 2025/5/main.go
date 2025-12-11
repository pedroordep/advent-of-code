package main

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/pedroordep/advent-of-code/utils"
)

type Range struct{ start, end int64 }

func main() {
	input := utils.GetInputFile(2025, 5)

	fmt.Println("Part 1", Part1(strings.Trim(string(input), "\n")))
	fmt.Println("Part 2", Part2(strings.Trim(string(input), "\n")))
}

func Part1(input string) int64 {
	parts := strings.Split(input, "\n\n")
	var ingredientsRange []Range

	for _, ingRange := range strings.Split(parts[0], "\n") {
		rangeSplit := strings.Split(ingRange, "-")
		start, _ := strconv.ParseInt(rangeSplit[0], 10, 64)
		end, _ := strconv.ParseInt(rangeSplit[1], 10, 64)
		ingredientsRange = append(ingredientsRange, Range{
			start: start,
			end:   end,
		})
	}

	sum := int64(0)
	for _, ingreditent := range strings.Split(parts[1], "\n") {
		ing, _ := strconv.ParseInt(ingreditent, 10, 64)
		for _, r := range ingredientsRange {
			if ing >= r.start && ing <= r.end {
				// fmt.Println("Found ingredient", ing, "in range", r)
				sum++
				break
			}
		}
	}
	return sum
}

func Part2(input string) int64 {
	parts := strings.Split(input, "\n\n")
	var ranges []Range

	for _, line := range strings.Split(parts[0], "\n") {
		rangeSplit := strings.Split(line, "-")
		start, _ := strconv.ParseInt(rangeSplit[0], 10, 64)
		end, _ := strconv.ParseInt(rangeSplit[1], 10, 64)
		ranges = append(ranges, Range{start, end})
	}

	slices.SortFunc(ranges, func(a, b Range) int {
		return cmp.Compare(a.start, b.start)
	})

	merged := []Range{}
	for _, r := range ranges {
		n := len(merged)
		if n == 0 || merged[n-1].end < r.start-1 {
			merged = append(merged, r)
		} else {
			if r.end > merged[n-1].end {
				merged[n-1].end = r.end
			}
		}
	}

	// for _, r := range merged {
	// 	fmt.Println(r, r.end-r.start+1)
	// }

	result := int64(0)
	for _, r := range merged {
		result += r.end - r.start + 1
	}
	return result
}
