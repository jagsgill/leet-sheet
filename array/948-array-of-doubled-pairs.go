package array

import (
	"sort"
)

func canReorderDoubled(A []int) bool {
	// count each value
	counts := make(map[int]int, 512)
	for i := 0; i < len(A); i++ {
		counts[A[i]]++
	}
	// we will scan the keys in increasing order, looking at whether enough pairings can be made
	// for the current elem i and its double, 2i

	// order the keys
	orderedKeys := make([]int, 0, len(counts))
	for k := range counts {
		orderedKeys = append(orderedKeys, k)
	}
	sort.Ints(orderedKeys)
	// orderedKeys is now ascending sorted

	var pairedIdx int
	for _, i := range orderedKeys {
		if counts[i] == 0 {
			continue
		}
		if i < 0 {
			if i & 0x1 == 1 {
				return false
			}
			pairedIdx = i / 2
		} else if i > 0 {
			pairedIdx = 2 * i
		} else {
			if counts[i] & 0x1 == 1 { // odd # of zeros can't be paired
				return false
			}
			counts[i] = 0
			continue
		}
		iCount := counts[i]
		pairedCnt := counts[pairedIdx]
		if iCount > pairedCnt {
			return false
		} else {
			counts[i] = 0
			counts[pairedIdx] = pairedCnt - iCount
		}
	}
	// counts is now zeroed since all elements were paired
	return true
}
