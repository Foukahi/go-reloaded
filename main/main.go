package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"stock"
)

var (
	Danger string = "❌ \033[31m"
	Nprobl string = "❗ \033[33m"
	ok     string = "\033[32m"
)

func main() {
	args := os.Args[1:]

	inFile, outFile, decision := "", "", ""

	if len(args) == 0 {
		fmt.Println(Danger + "File text and file result missing")
		return
	}
	if len(args) == 1 {

		fmt.Print(Nprobl + "Output file's name is missing\nWould you like to choose Output file's name? [y/n]: ")
		fmt.Scanf("%s", &decision)
		if decision == "y" {
			fmt.Print(ok + "Insert the file name: \033[37m")
			fmt.Scanf("%s", &outFile)
		} else if decision == "n" {
			outFile = "Def_result.txt"
			fmt.Println(Nprobl+"The result will be placed in", outFile, "file by default")
		}

	} else {
		outFile = args[1]
	}
	if len(args) > 2 {
		fmt.Println(Danger + "Too many files :(")
		return
	}
	inFile = args[0]

	textFile, err := os.Open(inFile)
	if err != nil {
		fmt.Println(Danger + "Something wrong check your files :)")
		return
	}

	resFile, err := os.Create(outFile)
	if err != nil {
		fmt.Println(Danger + "Something wrong check your files :)")
		return
	}

	scanner := bufio.NewScanner(textFile)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			resFile.WriteString("\n")
			continue
		}
		stock.Tab = strings.Fields(string(line))

		text := stock.Modifie()
		resFile.WriteString(text + "\n")
	}
}
