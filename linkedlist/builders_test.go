package linkedlist

import (
	"reflect"
	"testing"
)

func TestBuildListNode(t *testing.T) {
	spec := []int{1, 2, 3}
	var head = BuildListNode(spec)
	var actual = make([]int, 0, 3)
	for curr := head; curr != nil; curr = curr.Next {
		actual = append(actual, curr.Val)
	}
	if !reflect.DeepEqual(spec, actual){
		t.Errorf("expected %v, but got %v", spec, actual)
	}
}

func TestListNodeToArray(t *testing.T) {
	tail := &ListNode{3, nil}
	mid := &ListNode{2, tail}
	head := &ListNode{1, mid}

	expected := []int{1, 2, 3}
	actual := ListNodeToArray(head)

	if !reflect.DeepEqual(expected, actual){
		t.Errorf("expected %v, but got %v", expected, actual)
	}
}
