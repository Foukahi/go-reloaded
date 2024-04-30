package stock

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	Tab      []string
	Reminder string = "\033[33mMake sure your options are write"
	blocker  string = "\033[31mYou have a problem in your options"
)

type transformation func(string) string

var transformations = map[string]transformation{
	"(hex)": HexToDec,
	"(bin)": BinToDec,
	"(cap)": Cap,
	"(cap,": Cap,
	"(low)": strings.ToLower,
	"(low,": strings.ToLower,
	"(up)":  strings.ToUpper,
	"(up,":  strings.ToUpper,
}

func Modifie() string {
	if _, ok := transformations[Tab[0]]; ok {
		Tab[0] = ""
		if len(Tab) != 1 && Tab[1][len(Tab[1])-1:] == ")" {
			Tab[1] = ""
		}
		fmt.Println(Reminder)
	}
	for i := 1; i < len(Tab); i++ {
		if len(Tab[i]) != 0 && Tab[i][len(Tab[i])-1:] == "," && i != len(Tab)-1 && Tab[i+1][len(Tab[i+1])-1:] == ")" {
			if transform, ok := transformations[Tab[i]]; ok {

				num, err := strconv.Atoi(Tab[i+1][0 : len(Tab[i+1])-1])
				if err != nil {
					fmt.Println(blocker)
					break
				}
				if num > i {
					Tab = append(Tab[:i], Tab[i+2:]...)
					i--
					fmt.Println(Reminder)
					continue
				}
				for j := 0; j <= num; j++ {
					Tab[i-j] = transform(Tab[i-j])
				}
				Tab = append(Tab[:i], Tab[i+2:]...)
				i--
			}
		} else if transform, ok := transformations[Tab[i]]; ok && Tab[i][len(Tab[i])-1:] == ")" {
			Tab[i-1] = transform(Tab[i-1])
			Tab = append(Tab[:i], Tab[i+1:]...)
			i--
		}
	}
	FnlTxt := JoinSp(Tab)
	return AdjustText(FnlTxt)
}
