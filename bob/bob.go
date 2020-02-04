package bob

import (
	"regexp"
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	answer := ""

	containsLetter, _ := regexp.MatchString("[A-Z]+", remark)

	if len(remark) == 0 {
		answer = "Fine. Be that way!"
	} else if remark == strings.ToUpper(remark) && containsLetter {
		answer = "Whoa, chill out!"
		if strings.Contains(remark, "?") {
			answer = "Calm down, I know what I'm doing!"
		}
	} else if strings.HasSuffix(remark, "?") {
		answer = "Sure."
	} else {
		answer = "Whatever."
	}
	return answer
}
