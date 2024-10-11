package code

import "fmt"

// TIP FindPeakElement find peak element
// <a href="https://leetcode-cn.com/problems/find-peak-element/">LeetCode 链接</a>
func findPeakElement(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	if nums[0] > nums[1] {
		return 0
	}
	if nums[n-1] > nums[n-2] {
		return n - 1
	}
	l, r, m, ans := 1, n-2, 0, -1
	for l <= r {
		m = l + (r-l)>>1
		if nums[m] > nums[m+1] {
			fmt.Println("m:", m)
			ans = m
			r = m - 1
			continue
		}
		l = m + 1
	}
	return ans
}
