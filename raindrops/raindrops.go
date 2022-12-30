// Package raindrops provides converters for raindrops
package raindrops

import (
	"strconv"
	"strings"
)

// Convert transforms raindrops to sounds
func Convert(input int) string {
	drops := []struct {
		number int
		sound  string
	}{
		{
			number: 3,
			sound:  "Pling",
		},
		{
			number: 5,
			sound:  "Plang",
		},
		{
			number: 7,
			sound:  "Plong",
		},
	}

	var res = make([]string, 0)
	for _, v := range drops {
		if input%v.number == 0 {
			res = append(res, v.sound)
		}
	}
	if len(res) != 0 {
		return strings.Join(res, "")
	}
	return strconv.Itoa(input)
}
