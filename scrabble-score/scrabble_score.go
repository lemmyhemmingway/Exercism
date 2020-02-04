package scrabble

import (
	"strings"
)

func Score(answer string) int {
	answer = strings.ToUpper(answer)
	result := 0
	a := "A, E, I, O, U, L, N, R, S, T" // 1
	d := "D, G"							// 2
	b := "B, C, M, P"					// 3
	f := "F, H, V, W, Y" 				// 4
	k := "K"							// 5
	j := "J, X"							// 8
	q := "Q, Z"							// 10
	for _, c := range answer {
		var ch string = string(c)
		if strings.Contains(a, ch) {
			result += 1
		} else if strings.Contains(d, ch) {
			result += 2
		} else if strings.Contains(b, ch) {
			result += 3
		} else if strings.Contains(f, ch) {
			result += 4
		} else if strings.Contains(k, ch) {
			result += 5
		} else if strings.Contains(j, ch) {
			result += 8
		} else if strings.Contains(q, ch) {
			result += 10
		}

	}
	return result
}

