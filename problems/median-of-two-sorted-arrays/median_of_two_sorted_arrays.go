package medianoftwosortedarrays

import (
	"slices"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merge := append(nums1, nums2...)

	slices.Sort(merge)

	middle := len(merge) / 2

	if len(merge)%2 == 0 {
		return float64(merge[middle-1]+merge[middle]) / 2.0
	} else {
		return float64(merge[middle])
	}
}
