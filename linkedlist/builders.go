package linkedlist

type ListNode struct {
	Val int
	Next *ListNode
}

func BuildListNode(spec []int) *ListNode {
	dummyHead := new(ListNode)
	defer func() {
		dummyHead.Next = nil
	}()

	curr := dummyHead
	for i := 0; i < len(spec); i++ {
		curr.Next = &ListNode{spec[i], nil}
		curr = curr.Next
	}
	return dummyHead.Next
}

func ListNodeToArray(head *ListNode) []int {
	arr := make([]int, 0)
	for curr := head; curr != nil; curr = curr.Next {
		arr = append(arr, curr.Val)
	}
	return arr
}

