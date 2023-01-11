package allergies

/*
 * A bit of a fantasy solution but here it goes.
 * From the allergen values we can see that value is 2^orderOfAllergen
 *
 * So we can treat the total score as a binary representation of the enums
 * If allergen is included than that bit is 1 otherwise 0.
 *
 * From this perspective:
 * all we need to do is check if the final score has the bit for the Allergen flipped to 1.
 *
 * This can be done by pushing a bit, ordinal number of times to the left and then bitwise-and it with the score.
 * The result being zero will mean that this allergen is not in the list.
 * // Pushing a bit ordinal number of times is non-coincidentally is the value of the allergen.
 *
 *
 * Example :
 * Score 7
 *
 * Binary representation -> 00000111
 * Check if allergic to tomatoes.
 * Tomatoes ordinal -> 4
 * push 1 to the left 4 times -> 10000 (makes it 16, the value of allergen)
 *
 * bitwise & -> score & tomato -> if the result is zero than no tomato. if nonzero than tomato.
 * 00000111 & 10000 = 00000000 -> not included
 * This solution can be potentially optimized by using an enum instead if a slice.
 * */

func getList() []string {
	return []string{
		"eggs",
		"peanuts",
		"shellfish",
		"strawberries",
		"tomatoes",
		"chocolate",
		"pollen",
		"cats"}
}

func getIndex(al string) int {
	for i, v := range getList() {
		if v == al {
			return i
		}
	}
	return -1
}

func Allergies(allergies uint) []string {
	var res []string

	for _, v := range getList() {
		if AllergicTo(allergies, v) {
			res = append(res, v)
		}
	}

	return res
}

func AllergicTo(allergies uint, allergen string) bool {
	return allergies&(1<<getIndex(allergen)) != 0
}
