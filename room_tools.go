package stock

import (
	"fmt"
	"os"
	"strconv"
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
	res := ""
	ch := []rune(s)

	if IsString(string(ch[0])) && IsLower(string(ch[0])) {
		ch[0] = ch[0] - 32
	}
	for i := 0; i < len(ch); i++ {
		res += string(ch[i])
	}

	return res
}

func IsString(s string) bool {
	res := false
	str := []rune(s)

	for i := 0; i < len(str); i++ {
		if (str[i] >= 'a' && str[i] <= 'z') || (str[i] >= 'A' && str[i] <= 'Z') {
			res = true
		} else {
			res = false
			break
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
