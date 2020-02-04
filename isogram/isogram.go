package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(str string) bool {
	str = strings.ToLower(str)
	for _, ch := range str {
		if unicode.IsLetter(ch) && strings.Count(str, string(ch)) > 1 {
			return false
		}

	}
	return true
}
