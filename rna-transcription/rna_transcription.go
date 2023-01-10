package strand

import "strings"

func ToRNA(dna string) string {
	if len(dna) == 0 {
		return ""
	}

	transform := map[string]string{
		"G": "C",
		"C": "G",
		"T": "A",
		"A": "U",
	}

	dnaSlice := strings.Split(dna, "")
	res := make([]string, 0)
	for _, d := range dnaSlice {
		c, _ := transform[d]
		res = append(res, c)
	}

	return strings.Join(res, "")
}
