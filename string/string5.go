package string

/*
Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

Note: For the purpose of this problem, we define empty string as valid palindrome.

Example 1:

Input: "A man, a plan, a canal: Panama"
Output: true
Example 2:

Input: "race a car"
Output: false


Constraints:

s consists only of printable ASCII characters.
*/

func check(sr []int, start, end int) bool {
	for start < end {
		if sr[start] != sr[end] {
			return false
		}
		start++
		end--
	}

	return true
}

// IsPalindrome check Palindrom ( alphanumeric )
func IsPalindrome(s string) bool {

	r := make([]int, len(s))
	index := 0
	for _, v := range []rune(s) {
		n := int(v)
		if n >= 65 && n <= 90 {
			r[index] = n + 32
			index++

		} else if n >= 97 && n <= 122 {
			r[index] = n
			index++
		} else if n >= 48 && n <= 57 {
			r[index] = n
			index++
		}
	}
	return check(r[:index], 0, index-1)
}
