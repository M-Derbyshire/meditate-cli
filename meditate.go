package main

import (
	"fmt"
	"os"

	"github.com/M-Derbyshire/meditate-cli/help"
)

func main() {

	args := os.Args[1:]

	if len(args) > 0 {

		if args[0] == "help" {
			helpText := help.GetHelpText()
			fmt.Println(helpText)
		}

	}

}
