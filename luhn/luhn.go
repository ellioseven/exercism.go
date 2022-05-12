package luhn

import (
	"strings"
	"unicode"
)

func Valid(id string) bool {
	total := 0
	numbers := []rune(strings.ReplaceAll(id, " ", ""))

	if len(numbers) <= 1 {
		return false
	}

	for i := len(numbers); i != 0; i-- {
		s := numbers[i-1]

		if !unicode.IsDigit(rune(s)) {
			return false
		}

		// @url https://stackoverflow.com/questions/21322173/convert-rune-to-int
		number := s - '0'
		isEven := (len(numbers)-1-i)%2 == 0

		if isEven {
			dbl := number * 2

			if dbl >= 10 {
				total += int(dbl) - 9
			} else {
				total += int(dbl)
			}
		} else {
			total += int(number)
		}
	}

	return total%10 == 0
}
