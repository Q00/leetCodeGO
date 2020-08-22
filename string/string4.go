package string

// IsAnagram find anagram
/*
Given two strings s and t , write a function to determine if t is an anagram of s.

Example 1:

Input: s = "anagram", t = "nagaram"
Output: true
Example 2:

Input: s = "rat", t = "car"
Output: false
Note:
You may assume the string contains only lowercase alphabets.

Follow up:
What if the inputs contain unicode characters? How would you adapt your solution to such case?
*/
func IsAnagram(s string, t string) bool {
	smap := make([]int, 26)

	if len(s) != len(t) {
		return false
	}

	for _, v := range s {
		smap[v-'a']++
	}

	for _, v := range t {
		if smap[v-'a'] == 0 {
			return false
		}
		smap[v-'a']--
	}

	return true
}
