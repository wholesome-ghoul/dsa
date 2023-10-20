package non_overlapping_intervals

import "sort"

func EraseOverlapIntervals(intervals [][]int) int {
	return eraseOverlapIntervals(intervals)
}

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	intervalsToRemove := 0

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	left := 0
	right := 1
	for right < len(intervals) {
		if intervals[left][1] > intervals[right][0] {
			intervalsToRemove++
			right++
		} else {
			left = right
			right++
		}
	}

	return intervalsToRemove
}
