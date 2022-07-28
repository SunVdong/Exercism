package twelve

import "strings"

var (
	days = []string{
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
	gifts = []string{
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
)

func Verse(i int) string {
	i -= 1
	verse := "On the " + days[i] + " day of Christmas my true love gave to me: "
	if i > 0 {
		for j := i; j > 0; j-- {
			verse += gifts[j] + ", "
		}
		verse += "and "
	}

	return verse + gifts[0] + "."
}

func Song() string {
	var verses []string
	for i, _ := range days {
		verses = append(verses, Verse(i+1))
	}
	return strings.Join(verses, "\n")
}
