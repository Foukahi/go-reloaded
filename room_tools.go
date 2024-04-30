package stock

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func HexToDec(s string) string {
	DecHex, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		fmt.Println(blocker)
		os.Exit(0)
	}
	return strconv.Itoa(int(DecHex))
}

func BinToDec(s string) string {
	DecBin, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		fmt.Println(blocker)
		os.Exit(0)
	}
	return strconv.Itoa(int(DecBin))
}

func Cap(s string) string {
	s = strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
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

func JoinSp(tab []string) string {
	res := ""
	for i := 0; i < len(tab); i++ {
		if len(tab[i]) == 0 {
			continue
		} else {
			if i == len(tab)-1 {
				res += tab[i]
				break
			} else {
				res += tab[i] + " "
			}
		}
	}
	return res
}
