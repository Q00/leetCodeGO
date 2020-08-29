package string

import (
	"math"
	"strconv"
	"strings"
	"unicode"
)

func myAtoi(str string) int {
	result := ""

	s := strings.Trim(str, " ")

	index := 0

	if len(s) != 0 && s[0] != '-' && s[0] != '+' && !unicode.IsDigit(rune(s[0])) {
		return 0
	} else {
		for _, value := range s {
			if index == 0 && (value == '-' || value == '+') {
				result += string(value)
				continue
			} else {
				if unicode.IsDigit(rune(value)) {
					result += string(value)
				} else {
					break
				}
			}
			index += 1
		}
	}

	if len(s) == 0 {
		return 0
	} else {

		r, err := strconv.ParseInt(result, 10, 32)

		if err != nil {
			if index > 1 {
				if result[0] == '-' {
					return math.MinInt32
				} else {
					return math.MaxInt32
				}
			} else {
				return 0
			}
		}

		return int(r)
	}

	return 0
}
