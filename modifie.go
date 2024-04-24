package stock

import (
	"strconv"
	"strings"
)

func Modifie(txt string) string {
	tab := SplitWhiteSpaces(txt)
	ConvTxt := ""

	for i := 0; i < len(tab); i++ {

		if tab[i] == "(hex)" && i != 0 {

			DecHex, err := strconv.ParseInt(tab[i-1], 16, 64)
			if err != nil {
				continue
			}
			tab[i-1] = strconv.Itoa(int(DecHex))
			tab[i] = ""

		} else if tab[i] == "(bin)" && i != 0 {

			DecBin, err := strconv.ParseInt(tab[i-1], 2, 64)
			if err != nil {
				continue
			}
			tab[i-1] = strconv.Itoa(int(DecBin))
			tab[i] = ""

		} else if tab[i] == "(cap)" && i != 0 {

			tab[i-1] = Capitalize(tab[i-1])
			tab[i] = ""

		} else if tab[i] == "(cap," && i != 0 {

			num, err := strconv.Atoi(tab[i+1][0 : len(tab[i+1])-1])
			if err != nil || num > i || tab[i+1][len(tab[i+1])-1:] != ")" {
				continue
			} else {
				for j := 1; j <= num; j++ {
					tab[i-j] = Capitalize(tab[i-j])
				}
				tab[i] = ""
				tab[i+1] = ""
			}

		} else if tab[i] == "(low)" && i != 0 {

			tab[i-1] = strings.ToLower(tab[i-1])
			tab[i] = ""

		} else if tab[i] == "(low," && i != 0 {

			num, err := strconv.Atoi(tab[i+1][0 : len(tab[i+1])-1])
			if err != nil || num > i || tab[i+1][len(tab[i+1])-1:] != ")" {
				continue
			} else {
				for j := 1; j <= num; j++ {
					tab[i-j] = strings.ToLower(tab[i-j])
				}
				tab[i] = ""
				tab[i+1] = ""
			}

		} else if tab[i] == "(up)" && i != 0 {

			tab[i-1] = strings.ToUpper(tab[i-1])
			tab[i] = ""

		} else if tab[i] == "(up," && i != 0 {

			num, err := strconv.Atoi(tab[i+1][0 : len(tab[i+1])-1])
			if err != nil || num > i || tab[i+1][len(tab[i+1])-1:] != ")" {
				continue
			} else {
				for j := 1; j <= num; j++ {
					tab[i-j] = strings.ToUpper(tab[i-j])
				}
				tab[i] = ""
				tab[i+1] = ""
			}
		}
	}

	for i := 0; i < len(tab); i++ {
		if len(tab[i]) == 0 {
			continue
		} else {
			if i == len(tab)-1 {
				ConvTxt += tab[i]
				break
			} else {
				ConvTxt += tab[i] + " "
			}
		}
	}

	FnlTxt := AdjustText(CorrectAn(ConvTxt))

	return FnlTxt
}
