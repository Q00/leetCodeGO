/*
Given a non-empty array of integers, every element appears twice except for one. Find that single one.

Note:

Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

Example 1:

Input: [2,2,1]
Output: 1
Example 2:

Input: [4,1,2,1,2]
Output: 4
*/

package array

import (
	"fmt"
)

// SingleNumber single number
func SingleNumber(nums []int) int {
	save := map[int]bool{}
	for index, value := range nums {
		if save[nums[index]] != true {
			save[value] = true
		} else {
			save[value] = false
		}
	}
	single := 0

	for index, value := range save {
		if value == true {
			single = index
		}
	}
	fmt.Println(save)

	return single
}
