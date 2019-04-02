package house

import (
	"bytes"
)

var texts = []string{
	"This is ",
	"the house that Jack built.",
	"the malt\nthat lay in ",
	"the rat\nthat ate ",
	"the cat\nthat killed ",
	"the dog\nthat worried ",
	"the cow with the crumpled horn\nthat tossed ",
	"the maiden all forlorn\nthat milked ",
	"the man all tattered and torn\nthat kissed ",
	"the priest all shaven and shorn\nthat married ",
	"the rooster that crowed in the morn\nthat woke ",
	"the farmer sowing his corn\nthat kept ",
	"the horse and the hound and the horn\nthat belonged to ",
}

// Verse ...
func Verse(v int) string {
	return texts[0] + verse(v)
}

func verse(v int) string {
	if v == 1 {
		return texts[1]
	}

	return texts[v] + verse(v-1)
}

// Song ...
func Song() string {

	buf := bytes.Buffer{}
	buf.WriteString(Verse(1))

	for v := 2; v <= 12; v++ {
		buf.WriteString("\n\n")
		buf.WriteString(Verse(v))
	}

	return buf.String()
}
