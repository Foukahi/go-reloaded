package stock

import (
	"strings"
)

func CorrectAn(s string) string {
	tab := strings.Split(s, " ")
	res := ""

	for i := 0; i < len(tab); i++ {
		if tab[i] == "a" && i < len(tab)-1 {
			nxt := tab[i+1]
			if len(nxt) != 0 && (nxt[0] == 'a' || nxt[0] == 'e' || nxt[0] == 'i' || nxt[0] == 'o' || nxt[0] == 'u' || nxt[0] == 'h' || nxt[0] == 'A' || nxt[0] == 'E' || nxt[0] == 'I' || nxt[0] == 'O' || nxt[0] == 'U' || nxt[0] == 'H') {
				tab[i] = "an"
			}
		}
		if tab[i] == "A" && i < len(tab)-1 {
			nxt := tab[i+1]
			if len(nxt) != 0 && (nxt[0] == 'a' || nxt[0] == 'e' || nxt[0] == 'i' || nxt[0] == 'o' || nxt[0] == 'u' || nxt[0] == 'h' || nxt[0] == 'A' || nxt[0] == 'E' || nxt[0] == 'I' || nxt[0] == 'O' || nxt[0] == 'U' || nxt[0] == 'H') {
				tab[i] = "AN"
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

func IsLower(s string) bool {
	res := false
	str := []rune(s)

	for i := 0; i < len(str); i++ {
		if str[i] >= 'a' && str[i] <= 'z' {
			res = true
		} else {
			res = false
			break
		}
	}
	return res
}

func Capitalize(s string) string {
	res := ""
	ch := []rune(s)

	if len(ch) != 0 || IsLower(string(ch[0])) {
		ch[0] = ch[0] - 32
	}
	for i := 0; i < len(ch); i++ {
		res += string(ch[i])
	}

	return res
}
