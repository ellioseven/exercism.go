package twelve

import (
	"fmt"
	"strings"
)

func getDayLabels() map[int]string {
	return map[int]string{
		1:  "first",
		2:  "second",
		3:  "third",
		4:  "fourth",
		5:  "fifth",
		6:  "sixth",
		7:  "seventh",
		8:  "eighth",
		9:  "ninth",
		10: "tenth",
		11: "eleventh",
		12: "twelfth",
	}
}

func getNumLabels() map[int]string {
	return map[int]string{
		1:  "a",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
		11: "eleven",
		12: "twelve",
	}
}

func getItems() map[int]string {
	return map[int]string{
		1:  "Partridge in a Pear Tree",
		2:  "Turtle Doves",
		3:  "French Hens",
		4:  "Calling Birds",
		5:  "Gold Rings",
		6:  "Geese-a-Laying",
		7:  "Swans-a-Swimming",
		8:  "Maids-a-Milking",
		9:  "Ladies Dancing",
		10: "Lords-a-Leaping",
		11: "Pipers Piping",
		12: "Drummers Drumming",
	}
}

func Verse(i int) string {
	dayLabels := getDayLabels()
	numLabels := getNumLabels()
	items := getItems()

	sb := ""
	for j := 1; j <= i; j++ {
		if j == 1 && i == 1 {
			sb = fmt.Sprintf("%s %s", numLabels[j], items[j])
		} else if j == 1 {
			sb = fmt.Sprintf("and %s %s", numLabels[j], items[j])
		} else {
			sb = fmt.Sprintf("%s %s, %s", numLabels[j], items[j], sb)
		}
	}

	return fmt.Sprintf(
		"On the %s day of Christmas my true love gave to me: %s.",
		dayLabels[i],
		sb,
	)
}

func Song() string {
	const LENGTH = 12
	sb := strings.Builder{}

	for i := 1; i <= LENGTH; i++ {
		v := Verse(i)
		if i == LENGTH {
			sb.WriteString(v)
		} else {
			sb.WriteString(v + "\n")
		}
	}

	return sb.String()
}
