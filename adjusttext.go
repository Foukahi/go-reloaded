package stock

import "regexp"

func AdjustText(s string) string {
	s = regexp.MustCompile(`\s*([.,!?/:;])`).ReplaceAllString(s, "$1")
	s = regexp.MustCompile(`([.,!?/:;])\s*([^.,!?/:;])`).ReplaceAllString(s, "$1 $2")

	s = regexp.MustCompile(`'\s+([^']*)\s'`).ReplaceAllString(s, "'${1}'")

	return s
}

func SplitWhiteSpaces(s string) []string {
	var res []string
	ind := -1
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == '\n' {
			if ind != -1 {
				res = append(res, s[ind:i])
				ind = -1
			}
		} else if ind == -1 {
			ind = i
		}
	}
	if ind != -1 {
		res = append(res, s[ind:])
	}
	return res
}
