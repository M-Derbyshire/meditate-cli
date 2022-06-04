package commands

import (
	"errors"
	"sort"
	"strings"

	"github.com/M-Derbyshire/meditate-cli/listFile"
)

// Get the help text for the application
func Help() string {
	t := `Meditate CLI - V1.0.0
	
Meditate CLI allows you to create a list of words/concepts, that the application can then choose from for you at a later time. The idea is that you can recieve a random word/concept to meditate, reflect, and/or pray on.

Commands:

add <string> - Add a string to your list
remove <string> - Remove a string from your list
list - Lists the whole list, in alphabetical order
search <string> - Search the list for any item containing the given string
No arguments - Returns a randomly selected item from your list. The randomness of the choice is actually weighted, so items that have been chosen recently are less likely to be selected.

Created by Matthew Stuart Derbyshire - md-developer.uk`

	return t
}

// Get the full list (as a single string), in alphabetical order, seperated by new-lines.
func List(path string) (string, error) {

	list, err := listFile.LoadListFromFile(path)
	if err != nil {
		return "", errors.New("Error while loading list: " + err.Error())
	}

	sort.Strings(list)

	return strings.Join(list, "\n"), nil
}
