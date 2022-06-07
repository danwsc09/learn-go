package main

import (
	"fmt"
)

func main() {
	nums1 := []int{2, 7, 10, 15, 16, 0, 0, 0, 0, 0}
	nums2 := []int{3, 4, 9, 11, 12}
	merge(nums1, 5, nums2, 5)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
		return
	}

	p1, p2, write := m-1, n-1, m+n-1
	for p1 > -1 && p2 > -1 {
		if nums2[p2] > nums1[p1] {
			nums1[write] = nums2[p2]
			p2--
		} else {
			nums1[write] = nums1[p1]
			p1--
		}
		write--
	}

	for p1 > -1 {
		nums1[write] = nums1[p1]
		p1--
		write--
	}

	for p2 > -1 {
		nums1[write] = nums2[p2]
		p2--
		write--
	}
}
