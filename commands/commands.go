package commands

import (
	"errors"
	"math/rand"
	"sort"
	"strings"
	"time"

	"github.com/M-Derbyshire/meditate-cli/listFile"
	"github.com/M-Derbyshire/meditate-cli/strList"
)

// Get the help text for the application
func Help() string {
	t := `Meditate CLI - V1.0.0
	
Meditate CLI allows you to create a list of words/concepts that the application can then choose from for you at a later time. The idea is that you can recieve a random word/concept to meditate, reflect, and/or pray on.

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

// Get any items from the list (as a single string), that contain the given substring. Values will be ordered by length, and seperated with line breaks
func Search(listFilePath, substringToFind string) (string, error) {

	fullList, err := listFile.LoadListFromFile(listFilePath)
	if err != nil {
		return "", errors.New("Error while loading list: " + err.Error())
	}

	results := strList.FindBySubstring(fullList, substringToFind)
	sort.SliceStable(results, func(i, j int) bool {
		return len(fullList[i]) < len(fullList[j])
	})

	return strings.Join(results, "\n"), nil
}

// Add an item to the list. Result text will be a message saying that the item was added.
// If the item already exists (case insensitive), an error will be returned stating that
func Add(listFilePath, newItem string) (string, error) {

	currentList, loadErr := listFile.LoadListFromFile(listFilePath)
	if loadErr != nil {
		return "", errors.New("Error while loading list: " + loadErr.Error())
	}

	if strList.Contains(currentList, newItem) {
		return "", errors.New("The given item already exists in the list")
	}

	writeErr := listFile.AppendToListInFile(listFilePath, []string{newItem})
	if writeErr != nil {
		return "", errors.New("Error while adding item to list: " + writeErr.Error())
	}

	return "Item added to list", nil

}

// Remove an item from the list (case sensitive). Result text will be a message saying that the item was removed.
// If the item isn't found, an error will be returned stating that.
// This will only remove the first instance, but we don't allow duplicates
func Remove(listFilePath, itemToRemove string) (string, error) {

	list, loadErr := listFile.LoadListFromFile(listFilePath)
	if loadErr != nil {
		return "", errors.New("Error while loading list: " + loadErr.Error())
	}

	listOriginalLen := len(list)

	list = strList.RemoveFirstInstance(list, itemToRemove)
	if len(list) >= listOriginalLen {
		return "", errors.New("Item was not found in list")
	}

	writeErr := listFile.ReplaceListInFile(listFilePath, list)
	if writeErr != nil {
		return "", errors.New("Error while saving list: " + writeErr.Error())
	}

	return "Item removed from list", nil
}

// Pick an item from the list. This is a random choice, but the randomness is "weighted" in a way that favours items near the top of the list.
// If very few items are in the list, the one at the top of the list will be returned.
// Once an item is chosen, it will be moved to the bottom of the list.
func Choose(listFilePath string) (string, error) {

	var choice string
	var maxChoiceNum uint16 // To be compared against a random number, to determine if an item should be chosen

	list, loadErr := listFile.LoadListFromFile(listFilePath)
	if loadErr != nil {
		return "", errors.New("Error while loading list: " + loadErr.Error())
	}

	if len(list) == 0 {
		return "The list is currently empty", nil
	}

	if len(list) < 10 {
		choice = list[0]
	}

	if choice == "" {
		//We're going to loop through the items, running a random number generator on each one to decide whether or not that should be picked.
		//The likelyhood that an item will be picked is determined by how many items there are.

		maxChoiceNum = uint16((100 / len(list)) * 16)
		rand.Seed(time.Now().UnixNano())

		//We only want to loop through the list a certain amount of times. If we haven't chosen anything after that, we'll just choose the first item
	choiceAttemptLoop:
		for choiceAttempts := 0; choiceAttempts < 5; choiceAttempts++ {

			for _, currentItem := range list {
				if uint16(rand.Intn(100)) <= maxChoiceNum {
					choice = currentItem
					break choiceAttemptLoop
				}
			}

		}

		if choice == "" {
			choice = list[0]
		}
	}

	list = strList.MoveToEnd(list, choice)
	saveErr := listFile.ReplaceListInFile(listFilePath, list)
	if saveErr != nil {
		return "", errors.New("Error while saving list: " + loadErr.Error())
	}

	return "Meditate on: " + choice, nil
}
