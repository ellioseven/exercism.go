package grains

import (
	"errors"
)

func Square(number int) (uint64, error) {
	if number <= 0 || number >= 65 {
		return 0, errors.New("passed zero")
	}

	var total uint64 = 0

	// many solutions use bit operators to achieve the same result, which is
	// performant, however for readability I have chosen a for loop.
	for i := 0; i < number; i++ {
		if i == 0 {
			total += 1
		} else {
			total = total * 2
		}
	}

	return total, nil
}

func Total() uint64 {
	var total uint64
	return total - 1
}
