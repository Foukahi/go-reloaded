package stock

import (
	"regexp"
	"strings"
)

func AdjustText(s string) string {
	s = regexp.MustCompile(`\s*([.,!?/:;])`).ReplaceAllString(s, "$1")
	s = regexp.MustCompile(`([.,!?/:;])\s*([^.,!?/:;])`).ReplaceAllString(s, "$1 $2")
	return AdjustQuote(s)
}

func AdjustQuote(s string) string {
	tab := []rune(s)
	res := ""
	first := true
	last := false
	if len(tab) > 2 {
		if tab[0] == '\'' {
			first = false
			last = true
			if tab[1] == ' ' {
				tab = append(tab[:1], tab[2:]...)
			}
		}
		for i := 1; i < len(tab)-1; i++ {
			if first && !last && tab[i] == '\'' && (tab[i-1] == ' ' || tab[i+1] == ' ') {
				first = false
				if tab[i+1] == ' ' {
					tab = append(tab[:i+1], tab[i+2:]...)
				}
				last = true

			} else if last && !first && tab[i] == '\'' && (tab[i-1] == ' ' || tab[i+1] == ' ') {
				last = false
				if tab[i-1] == ' ' {
					tab = append(tab[:i-1], tab[i:]...)
				}
				first = true
			}
		}
		if tab[len(tab)-1] == '\'' && !first {
			if tab[len(tab)-2] == ' ' {
				tab = append(tab[:len(tab)-2], tab[len(tab)-1:]...)
			}
		}
	}
	for i := 0; i < len(tab); i++ {
		res += string(tab[i])
	}
	return CorrectAn(res)
}

func CorrectAn(s string) string {
	tab := strings.Split(s, " ")
	res := ""
	for i := 0; i < len(tab); i++ {
		if (tab[i] == "a" || tab[i] == "'a") && i < len(tab)-1 {
			nxt := tab[i+1]
			if len(nxt) != 0 && (nxt[0] == 'a' || nxt[0] == 'e' || nxt[0] == 'i' || nxt[0] == 'o' || nxt[0] == 'u' || nxt[0] == 'h' || nxt[0] == 'A' || nxt[0] == 'E' || nxt[0] == 'I' || nxt[0] == 'O' || nxt[0] == 'U' || nxt[0] == 'H') {
				if tab[i] == "'a" {
					tab[i] = "'an"
				} else {
					tab[i] = "an"
				}
			}
		}
		if (tab[i] == "A" || tab[i] == "'A") && i < len(tab)-1 {
			nxt := tab[i+1]
			if len(nxt) != 0 && (nxt[0] == 'a' || nxt[0] == 'e' || nxt[0] == 'i' || nxt[0] == 'o' || nxt[0] == 'u' || nxt[0] == 'h' || nxt[0] == 'A' || nxt[0] == 'E' || nxt[0] == 'I' || nxt[0] == 'O' || nxt[0] == 'U' || nxt[0] == 'H') {
				if tab[i] == "'a" {
					tab[i] = "'AN"
				} else {
					tab[i] = "AN"
				}
			}
		}
		if i == len(tab)-1 {
			res += tab[i]
			break
		} else {
			res += tab[i] + " "
		}
	}
	return res
}
