package stock

import "regexp"

func AdjustText(s string) string {
	s = regexp.MustCompile(`\s*([.,!?/:;])`).ReplaceAllString(s, "$1")
	s = regexp.MustCompile(`([.,!?/:;])\s*([^.,!?/:;])`).ReplaceAllString(s, "$1 $2")

	s = regexp.MustCompile(`'\s+([^']*)\s'`).ReplaceAllString(s, "'${1}'")
	s = regexp.MustCompile(`"\s([^']*)\s"`).ReplaceAllString(s, "\"${1}\"")

	return s
}
