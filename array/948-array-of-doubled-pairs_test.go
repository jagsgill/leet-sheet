package array

import "testing"

func TestCanReorderDoubled(t *testing.T) {
	assertCanReorder([]int{2,1,2,1,1,1,2,2}, true, t)
}

func TestCanReorderDoubledZeros(t *testing.T) {
	assertCanReorder([]int{0, 0, 0}, false, t)
}

func assertCanReorder(arr []int, expected bool, t *testing.T) {
	actual := canReorderDoubled(arr)
	if expected != actual {
		t.Errorf("expected '%v' for %v", expected, arr)
	}
}

