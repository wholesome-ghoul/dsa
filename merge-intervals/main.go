package merge_intervals

import (
	"math"
	"sort"
)

func Merge(intervals [][]int) [][]int {
	return merge(intervals)
}

func merge(intervals [][]int) [][]int {
	mergedIntervals := [][]int{}
	sortByStart := func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	}

	sort.Slice(intervals, sortByStart)

	currInterval := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if currInterval[1] >= intervals[i][0] {
			currInterval[1] = int(math.Max(float64(currInterval[1]), float64(intervals[i][1])))
			continue
		}

		mergedIntervals = append(mergedIntervals, currInterval)
		currInterval = intervals[i]
	}
	mergedIntervals = append(mergedIntervals, currInterval)

	return mergedIntervals
}
