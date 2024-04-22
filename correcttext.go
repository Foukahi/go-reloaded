package stock

import "strings"

func CorrectText(s string) string {
	tab := strings.Split(s, " ")
	res := ""

	for i := 0; i < len(tab); i++ {
		if tab[i] == "a" {
			nxt := tab[i+1]
			if nxt[0] == 'a' || nxt[0] == 'e' || nxt[0] == 'i' || nxt[0] == 'o' || nxt[0] == 'u' || nxt[0] == 'h' || (nxt[0] == 'A' || nxt[0] == 'E' || nxt[0] == 'I' || nxt[0] == 'O' || nxt[0] == 'U' || nxt[0] == 'H') {
				tab[i] = "an"
			}
		}
		if tab[i] == "A" {
			nxt := tab[i+1]
			if (nxt[0] == 'a' || nxt[0] == 'e' || nxt[0] == 'i' || nxt[0] == 'o' || nxt[0] == 'u' || nxt[0] == 'h') || (nxt[0] == 'A' || nxt[0] == 'E' || nxt[0] == 'I' || nxt[0] == 'O' || nxt[0] == 'U' || nxt[0] == 'H') {
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
