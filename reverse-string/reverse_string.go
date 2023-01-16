package reverse

func Reverse(input string) string {
	var runes []rune
	for _, r := range input {
		runes = append([]rune{r}, runes...)
	}

	return string(runes)
}
