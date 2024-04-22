package main

import (
	"fmt"
	"os"
	"stock"
)

func main() {
	files := os.Args[1:]

	if len(files) == 0 {
		fmt.Println("File text and file result missing")
		return
	}
	if len(files) == 1 {
		fmt.Println("File result missing")
		return
	}
	if len(files) > 2 {
		fmt.Println("Too many files :(")
		return
	}

	cont, err := os.ReadFile(files[0])
	if err != nil {
		fmt.Println("Something wrong check your files :)")
		return
	}

	text := stock.Modifie(string(cont))

	stock.CreatFile(text, files[1])

}
