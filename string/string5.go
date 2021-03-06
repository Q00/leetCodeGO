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

func check(sr []byte, start, end int) bool {
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

	r := make([]byte, len(s))
	index := 0

	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			r[index] = byte(v + 32)
			index++

		} else if v >= 'a' && v <= 'z' {
			r[index] = byte(v)
			index++
		} else if v >= '0' && v <= '9' {
			r[index] = byte(v)
			index++
		}
	}
	return check(r[:index], 0, index-1)
}
