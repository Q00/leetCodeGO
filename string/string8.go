package string

import "strconv"

func countAndSay(n int) string {

	return recu("1", n)
}

func recu(str string, n int) string {
	if n == 1 {
		return str
	}

	char := ' '
	result := ""
	count := 0

	for i, v := range str {
		if char == ' ' || char == v {
			count++
		} else {
			result += strconv.Itoa(count) + string(char)
			count = 1
		}

		char = v

		if i == len(str)-1 {
			result += strconv.Itoa(count) + string(char)
		}
	}

	return recu(result, n-1)
}
