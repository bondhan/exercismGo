package twelve

type text struct {
	numberString string
	thing        string
}

var texts = []text{
	{"first", "a Partridge in a Pear Tree."},
	{"second", "two Turtle Doves, and "},
	{"third", "three French Hens, "},
	{"fourth", "four Calling Birds, "},
	{"fifth", "five Gold Rings, "},
	{"sixth", "six Geese-a-Laying, "},
	{"seventh", "seven Swans-a-Swimming, "},
	{"eighth", "eight Maids-a-Milking, "},
	{"ninth", "nine Ladies Dancing, "},
	{"tenth", "ten Lords-a-Leaping, "},
	{"eleventh", "eleven Pipers Piping, "},
	{"twelfth", "twelve Drummers Drumming, "},
}

func Verse(verse int) string {
	text := "On the " + texts[verse-1].numberString + " day of Christmas my true love gave to me: "
	for i := verse - 1; i >= 0; i-- {
		text += texts[i].thing
	}

	return text
}

func Song() string {
	song := ""
	for index, _ := range texts {
		song += Verse(index+1) + "\n"
	}

	return song
}
