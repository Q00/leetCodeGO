package string

import (
	"sort"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}

	sort.SliceStable(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})

	le := strs[0]

	for _, j := range strs {
		for i := len(le); i >= 0; i-- {
			// prefix(j[:i]) is alwasy shorter than string
			if strings.HasPrefix(le, j[:i]) {
				le = j[:i]
				break
			}
		}
	}

	return le

}
