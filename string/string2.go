package string

import (
	"math"
	"strconv"
)

func reverseRune(target []rune, start, end int) {
	for start < end {
		target[start], target[end] = target[end], target[start]
		start++
		end--
	}
}

// Reverse Reverse Integer
/*
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:

Input: 123
Output: 321
Example 2:

Input: -123
Output: -321
Example 3:

Input: 120
Output: 21
Note:
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
*/
func Reverse(x int) int {

	sign := 1
	if x < 0 {
		sign = -1
		x *= -1
	}
	s := strconv.Itoa(x)
	r := []rune(s)
	reverseRune(r, 0, len(r)-1)
	result, _ := strconv.Atoi(string(r))
	result *= sign
	if result > int(math.MinInt32) && result < int(math.MaxInt32)-1 {
		return result
	}
	return 0
}

// func reverse(x int) int {

//         sign := 1
//         if x < 0 {
//             sign*=-1
//             x*=-1
//         }
//         s:= strconv.Itoa(x)
//         r:= ""
//         for i:= len(s)-1; i>=0;i-- {
//             r += string(s[i])
//         }
//         result,_:= strconv.Atoi(r)
//         result*=sign
//     if result > int(math.MinInt32) && result < int(math.MaxInt32)-1 {
//         return result
//     } else {
//         return 0
//     }
// }

// func reverse(x int) int {

//     var b strings.Builder
//     sign:=1
//     if x < 0 {
//         sign *= -1
//         x *= -1
//     }
//     s := strconv.Itoa(x)
//     // fmt.Println(s[2])
//     b.Grow(len(s)+2)
//     for i:= len(s)-1; i>=0;i-- {

//         fmt.Fprintf(&b, string(s[i]))
//     }

//     r,_:= strconv.Atoi(b.String())
//     if r < math.MinInt32 || r > math.MaxInt32-1 {
//         return 0
//     }

//     return r* sign
// }
