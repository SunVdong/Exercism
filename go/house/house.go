package house

import "strings"

var m = [13][2]string{
	{"", ""},
	{"house", "Jack built"},
	{"malt", "lay in"},
	{"rat", "ate"},
	{"cat", "killed"},
	{"dog", "worried"},
	{"cow with the crumpled horn", "tossed"},
	{"maiden all forlorn", "milked"},
	{"man all tattered and torn", "kissed"},
	{"priest all shaven and shorn", "married"},
	{"rooster that crowed in the morn", "woke"},
	{"farmer sowing his corn", "kept"},
	{"horse and the hound and the horn", "belonged to"},
}

func Verse(v int) string {
	var subF func(int) string
	subF = func(n int) string {
		if n == 1 {
			return "house that Jack built."
		} else {
			return m[n][0] + "\n" + "that " + m[n][1] + " the " + subF(n-1)
		}
	}

	return "This is the " + subF(v)
}

func Song() string {
	var sb strings.Builder
	for i := 1; i < 12; i++ {
		sb.WriteString(Verse(i))
		sb.WriteString("\n\n")
	}
	sb.WriteString(Verse(12))
	return sb.String()
}
