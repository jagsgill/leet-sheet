package linkedlist

import (
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	var orig = BuildListNode([]int{1, 2, 3, 4, 5})
	var actual = removeNthFromEnd(orig, 2)
	var expected = BuildListNode([]int{1, 2, 3, 5})
	assertEqualLists(expected, actual, t)
}

func assertEqualLists(expected *ListNode, actual *ListNode, t *testing.T) {
	expArr := ListNodeToArray(expected)
	actArr := ListNodeToArray(actual)
	if !reflect.DeepEqual(expArr, actArr) {
		t.Errorf("expected %v, but got %v", expArr, actArr)
	}
}


