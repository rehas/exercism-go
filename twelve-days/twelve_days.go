package twelve

import (
	"fmt"
)

var verseIntroFormat string = "On the %s day of Christmas my true love gave to me"

var verses map[int][]string = map[int][]string{
	1:  {"first", "a Partridge in a Pear Tree"},
	2:  {"second", "two Turtle Doves"},
	3:  {"third", "three French Hens"},
	4:  {"fourth", "four Calling Birds"},
	5:  {"fifth", "five Gold Rings"},
	6:  {"sixth", "six Geese-a-Laying"},
	7:  {"seventh", "seven Swans-a-Swimming"},
	8:  {"eighth", "eight Maids-a-Milking"},
	9:  {"ninth", "nine Ladies Dancing"},
	10: {"tenth", "ten Lords-a-Leaping"},
	11: {"eleventh", "eleven Pipers Piping"},
	12: {"twelfth", "twelve Drummers Drumming"},
}

func Verse(i int) string {
	verseLeft := fmt.Sprintf(verseIntroFormat, verses[i][0])

	verseRight := verses[1][1]
	for index := 2; index <= i; index++ {
		currentVerse := verses[index][1]
		if index == 2 {
			verseRight = fmt.Sprintf("%s, and %s", currentVerse, verseRight)
			continue
		}
		verseRight = fmt.Sprintf("%s, %s", currentVerse, verseRight)
	}

	return fmt.Sprintf("%s: %s.", verseLeft, verseRight)
}

func Song() string {
	var song string
	for i := 1; i <= len(verses); i++ {
		song = fmt.Sprintf("%s%s", song, fmt.Sprintf("%s\n", Verse(i)))
	}
	return song[:len(song)-1] //remove last trailing newline
}