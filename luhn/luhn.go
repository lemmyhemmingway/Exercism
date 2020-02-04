package luhn

import (
	"strconv"
	"strings"
)

// Valid validates number using the luhn formula
func Valid(input string) bool {
	input = strings.Replace(input, " ", "", -1)
	if len(input) <= 1 {
		return false
	}

	input = reverse(input)

	sum, err := calculate(input)
	if err != nil {
		return false
	}

	return sum%10 == 0

}

func reverse(input string) string {
	reversedWord := ""

	for i := len(input) - 1; i >= 0; i-- {
		reversedWord += string(input[i])
	}
	return reversedWord
}

func calculate(input string) (int, error) {
	sum := 0
	for i := 0; i < len(input); i++ {
		r, err := strconv.Atoi(string(input[i]))
		if err != nil {
			return -1, err
		}
		if i%2 != 0 {
			double := r * 2
			if double > 9 {
				double -= 9
			}
			sum += double
		} else {
			sum += r
		}

	}
	return sum, nil
}
