package array

import "fmt"

/*
Given a non-empty array of digits representing a non-negative integer, increment one to the integer.

The digits are stored such that the most significant digit is at the head of the list, and each element in the array contains a single digit.

You may assume the integer does not contain any leading zero, except the number 0 itself.

Example 1:

Input: [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Example 2:

Input: [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
*/

// PlusOne array plus
func PlusOne(digits []int) []int {
	next := 0
	val := 0

	for i := len(digits) - 1; i >= 0; i-- {
		if i == len(digits)-1 {
			next = 1
		}
		val = digits[i] + next
		digits[i] = val % 10
		next = val / 10

	}

	if next > 0 {
		digits = append([]int{next}, digits...)
		fmt.Println(digits)
	}
	fmt.Println(digits)
	return digits
}
