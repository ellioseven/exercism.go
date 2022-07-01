package twelve

import (
	"strings"
)

var gifts = []string{
	"a Partridge in a Pear Tree",
	"two Turtle Doves",
	"three French Hens",
	"four Calling Birds",
	"five Gold Rings",
	"six Geese-a-Laying",
	"seven Swans-a-Swimming",
	"eight Maids-a-Milking",
	"nine Ladies Dancing",
	"ten Lords-a-Leaping",
	"eleven Pipers Piping",
	"twelve Drummers Drumming",
}

var days = []string{
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

func Verse(i int) string {
	num := i - 1
	verse := "On the " + days[num] + " day of Christmas my true love gave to me: "

	if num > 0 {
		for i := num; i > 0; i-- {
			verse += gifts[i] + ", "
		}

		verse += "and "
	}

	return verse + gifts[0] + "."
}

func Song() string {
	var verses []string

	for i := range days {
		verses = append(verses, Verse(i+1))
	}

	return strings.Join(verses, "\n")
}
