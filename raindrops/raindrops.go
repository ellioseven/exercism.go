package raindrops

import (
	"fmt"
	"strings"
)

type translation struct {
	num  int
	word string
}

func Convert(number int) string {
	// a string is a primitive type, which means it is read-only, and every
	// manipulation of it will create a new string, `strings.Builder` is a
	// memory efficient mechanism for writing strings.
	// @url https://bit.ly/3MGp9eW
	var output strings.Builder

	translations := []translation{
		{3, "Pling"},
		{5, "Plang"},
		{7, "Plong"},
	}

	for _, trans := range translations {
		if number%trans.num == 0 {
			output.WriteString(trans.word)
		}
	}

	if output.Len() == 0 {
		return fmt.Sprintf("%d", number)
	}

	return output.String()
}
