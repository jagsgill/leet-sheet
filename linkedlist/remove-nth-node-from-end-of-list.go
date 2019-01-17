package linkedlist

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 1 <= n <= len(list)
	// graft a dummy node to the head to simplify handling when n=len(list)
	var dummyHead = &ListNode{Val: 0, Next: head}
	// advance to nth node from start, hold prev-nth at head
	curr := dummyHead
	for i := 0; i < n; i++ {
		curr = curr.Next
	}
	nPlusOneBack := dummyHead
	// loop to end of list, advance iterator and prev-nth in each iteration
	for ; curr.Next != nil; curr = curr.Next {
		nPlusOneBack = nPlusOneBack.Next
	}

	// after loop terminates, excise Nth node then graft the resulting lists
	nPlusOneBack.Next = nPlusOneBack.Next.Next
	return dummyHead.Next
}