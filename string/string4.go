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
	short := s
	long := t
	if len(s) > len(t) {
		short = t
		long = s
	}
	for _, v := range short {
		smap[v-'a']++
	}

	for _, v := range long {
		if smap[v-'a'] == 0 {
			return false
		}
		smap[v-'a']--
	}

	return true
}
