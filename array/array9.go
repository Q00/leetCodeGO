package array

/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

// TwoSum add two propery to make target
func TwoSum(nums []int, target int) []int {
	// check := map[int][]int{}

	// for i, v := range nums {
	// 	check[v] = append(check[v], i)

	// 	if len(check[target-v]) >= 1 {
	// 		// same value has two property in array
	// 		if v == target-v && len(check[v]) > 1 {
	// 			return check[v]
	// 		} else if v != target-v {
	// 			return []int{check[target-v][0], check[v][0]}
	// 		}
	// 	}
	// }

	check := map[int]int{}

	for i, v := range nums {
		curr, exist := check[target-v]
		if exist {
			return []int{i, curr}
		}
	}

	return nums
}
