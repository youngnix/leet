package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(nums ...int) *ListNode {
	list := &ListNode{}
	node := &list

	for i, n := range nums {
		(*node).Val = n

		node = &(*node).Next

		if i < len(nums)-1 {
			*node = &ListNode{}
		}
	}

	return list
}

func (l *ListNode) Array() []int {
	arr := []int{}
	node := l

	for node != nil {
		arr = append(arr, node.Val)

		node = node.Next
	}

	return arr
}

func (l *ListNode) Equals(list *ListNode) bool {
	for {
		if l.Val != list.Val {
			return false
		}

		if l.Next == nil && list.Next == nil {
			return true
		}

		if l.Next == nil || list.Next == nil {
			return false
		}

		l = l.Next
		list = list.Next
	}
}
