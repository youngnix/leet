package medianoftwosortedarrays

import (
	"fmt"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	cases := []struct {
		nums1 []int
		nums2 []int
		want  float64
	}{
		{
			nums1: []int{1, 3},
			nums2: []int{2},
			want:  2.0,
		},
		{
			nums1: []int{1, 2},
			nums2: []int{3, 4},
			want:  2.5,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%v %v", c.nums1, c.nums2), func(t *testing.T) {
			got := findMedianSortedArrays(c.nums1, c.nums2)
			if got != c.want {
				t.Errorf("expected %v, got %v", c.want, got)
			}
		})
	}
}
