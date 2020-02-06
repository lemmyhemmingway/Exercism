package acronym

import (
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	r := strings.NewReplacer("-", " ", "_", " ")
	s = r.Replace(s)
	acc := ""
	words := strings.Fields(s)
	for _, word := range words {
		acc += string(word[0])
	}
	return strings.ToUpper(acc)
}
