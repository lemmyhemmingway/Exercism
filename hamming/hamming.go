package hamming

import "errors"

func Distance(a, b string) (int, error) {
	hamm := 0
	if len(a) != len(b) {
		return 0, errors.New("Hola")
	}
	for index := 0; index < len(a); index++ {
		if a[index] != b[index] {
			hamm += 1
		}
	}
	return hamm, nil

}
