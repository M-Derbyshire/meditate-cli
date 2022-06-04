package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/M-Derbyshire/meditate-cli/commands"
)

func main() {

	var resultText string
	var err error

	args := os.Args[1:]

	//The list file needs to be in the same directory as the executable, not in the current console location
	executablePath, err := os.Executable()
	listFilePath := filepath.Join(filepath.Dir(executablePath), "meditate_list")
	// listFilePath = "test.txt"

	if err == nil { //If we have successfully loaded the file path for the list

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

			if args[0] == "add" {
				if len(args) < 2 {
					resultText = "Please provide an item to be added to the list"
				} else {
					resultText, err = commands.Add(listFilePath, args[1])
				}
			}

			if args[0] == "remove" {
				if len(args) < 2 {
					resultText = "Please provide an item to be removed from the list"
				} else {
					resultText, err = commands.Remove(listFilePath, args[1])
				}
			}

		}

	}

	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Print(resultText)
	}
}
