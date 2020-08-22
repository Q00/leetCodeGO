package string

/*
Given a string, find the first non-repeating character in it and return its index. If it doesn't exist, return -1.

Examples:

s = "leetcode"
return 0.

s = "loveleetcode"
return 2.


Note: You may assume the string contains only lowercase English letters.
*/

// firstUniqChar First Unique Character in a String
// func firstUniqChar(s string) int {
//     smap:= make(map[rune]int)

//     for _,v := range s {
//         smap[v]++
//     }

//     for i,v := range s {
//         if smap[v] == 1{
//             return i
//         }
//     }

//     return -1
// }

// FirstUniqChar find unique letter
func FirstUniqChar(s string) int {
	smap := make([]int, 26)

	for _, v := range s {
		smap[v-'a']++
	}

	for i, v := range s {
		if smap[v-'a'] == 1 {
			return i
		}
	}

	return -1
}
