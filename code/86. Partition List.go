package code

func partition(head *ListNode, x int) *ListNode {
	lessHead := &ListNode{}
	less := lessHead
	moreHead := &ListNode{}
	more := moreHead
	for head != nil {
		if head.Val < x {
			less.Next = head
			less = less.Next
		} else {
			more.Next = head
			more = more.Next
		}
		head = head.Next
	}
	more.Next = nil
	less.Next = moreHead.Next
	return lessHead.Next
}
