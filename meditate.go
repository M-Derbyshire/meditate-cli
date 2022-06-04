package main

import (
	"fmt"
	"os"

	"github.com/M-Derbyshire/meditate-cli/commands"
)

func main() {

	args := os.Args[1:]
	listFilePath := "test1.txt"
	var resultText string
	var err error

	if len(args) > 0 {

		if args[0] == "help" {
			resultText = commands.Help()
		}

		if args[0] == "list" {
			resultText, err = commands.List(listFilePath)
		}

		if args[0] == "search" {
			if len(args) < 2 {
				resultText = "Please provide a substring to be searched for"
			} else {
				resultText, err = commands.Search(listFilePath, args[1])
			}
		}

	}

	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Print(resultText)
	}
}
