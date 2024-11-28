package addtwonumbers

import (
	"fmt"
	"leet/list"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	testCases := []struct {
		l1, l2, want *list.ListNode
	}{
		{
			l1:   list.New(2, 4, 3),
			l2:   list.New(5, 6, 4),
			want: list.New(7, 0, 8),
		},
		{
			l1:   list.New(0),
			l2:   list.New(0),
			want: list.New(0),
		},
		{
			l1:   list.New(9, 9, 9, 9, 9, 9, 9),
			l2:   list.New(9, 9, 9, 9),
			want: list.New(8, 9, 9, 9, 0, 0, 0, 1),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%v + %v", tc.l1.Array(), tc.l2.Array()), func(t *testing.T) {
			got := addTwoNumbers(tc.l1, tc.l2)

			if !tc.want.Equals(got) {
				t.Errorf("wanted %v, got %v", tc.want.Array(), got.Array())
			}
		})
	}
}
