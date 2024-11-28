package addtwonumbers

import "leet/list"

func addTwoNumbers(l1, l2 *list.ListNode) *list.ListNode {
	l3 := &list.ListNode{}

	l1Node := l1
	l2Node := l2
	l3Node := l3

	for {
		if l1Node != nil {
			l3Node.Val += l1Node.Val
		}

		if l2Node != nil {
			l3Node.Val += l2Node.Val
		}

		if (l1Node != nil && l1Node.Next != nil) || (l2Node != nil && l2Node.Next != nil) {
			l3Node.Next = &list.ListNode{}
		}

		if l3Node.Val >= 10 {
			carry := l3Node.Val / 10
			l3Node.Val %= 10

			if l3Node.Next == nil {
				l3Node.Next = &list.ListNode{}
			}

			l3Node.Next.Val += carry
		}

		if l1Node != nil {
			l1Node = l1Node.Next
		}

		if l2Node != nil {
			l2Node = l2Node.Next
		}

		if l3Node != nil {
			l3Node = l3Node.Next
		}

		if l1Node == nil && l2Node == nil {
			break
		}
	}

	return l3
}
