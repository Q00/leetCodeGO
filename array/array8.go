package array

import (
	"sort"
)

/*
Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Example:

Input: [0,1,0,3,12]
Output: [1,3,12,0,0]
Note:

You must do this in-place without making a copy of the array.
Minimize the total number of operations.
*/

// MoveZeroes move zero backward
func MoveZeroes(nums []int) {
	check := false
	for _, value := range nums {
		if value == 0 {
			check = true
		}
	}

	if check {
		// not guaranted argument of list
		// sort.Slice(nums, func(_, j int) bool {
		// return nums[j] == 0
		// })

		// use SliceStable
		sort.SliceStable(nums, func(_, j int) bool {
			return nums[j] == 0
		})
	}
}
