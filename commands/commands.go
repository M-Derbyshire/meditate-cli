package commands

import (
	"errors"
	"math"
	"math/rand"
	"sort"
	"strings"
	"time"

	"github.com/M-Derbyshire/meditate-cli/listfile"
	"github.com/M-Derbyshire/meditate-cli/strlist"
)

// Help returns the help text for the application
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

// List returns the full list (as a single string), in alphabetical order, seperated by new-lines.
func List(path string) (string, error) {

	list, err := listfile.LoadListFromFile(path)
	if err != nil {
		return "", errors.New("Error while loading list: " + err.Error())
	}

	sort.SliceStable(list, func(i, j int) bool {
		return strings.ToLower(list[i]) < strings.ToLower(list[j])
	})

	return strings.Join(list, "\n"), nil
}

// Search returns any items from the list (as a single string), that contain the given substring. Values will be ordered by length, and seperated with line breaks
func Search(listfilePath, substringToFind string) (string, error) {

	fullList, err := listfile.LoadListFromFile(listfilePath)
	if err != nil {
		return "", errors.New("Error while loading list: " + err.Error())
	}

	results := strlist.FindBySubstring(fullList, substringToFind)
	sort.SliceStable(results, func(i, j int) bool {
		return len(fullList[i]) < len(fullList[j])
	})

	return strings.Join(results, "\n"), nil
}

// Add will add an item to the list. Result text will be a message saying that the item was added.
// If the item already exists (case insensitive), an error will be returned stating that
func Add(listfilePath, newItem string) (string, error) {

	currentList, loadErr := listfile.LoadListFromFile(listfilePath)
	if loadErr != nil {
		return "", errors.New("Error while loading list: " + loadErr.Error())
	}

	if strlist.Contains(currentList, newItem) {
		return "", errors.New("The given item already exists in the list")
	}

	writeErr := listfile.AppendToListInFile(listfilePath, []string{newItem})
	if writeErr != nil {
		return "", errors.New("Error while adding item to list: " + writeErr.Error())
	}

	return "Item added to list", nil

}

// Remove will remove an item from the list (case sensitive). Result text will be a message saying that the item was removed.
// If the item isn't found, an error will be returned stating that.
// This will only remove the first instance, but we don't allow duplicates
func Remove(listfilePath, itemToRemove string) (string, error) {

	list, loadErr := listfile.LoadListFromFile(listfilePath)
	if loadErr != nil {
		return "", errors.New("Error while loading list: " + loadErr.Error())
	}

	listOriginalLen := len(list)

	list = strlist.RemoveFirstInstance(list, itemToRemove)
	if len(list) >= listOriginalLen {
		return "", errors.New("Item was not found in list")
	}

	writeErr := listfile.ReplaceListInFile(listfilePath, list)
	if writeErr != nil {
		return "", errors.New("Error while saving list: " + writeErr.Error())
	}

	return "Item removed from list", nil
}

// Choose will pick an item from the list. This is a random choice, but the randomness is "weighted" in a way that favours items near the top of the list.
// If very few items are in the list, the one at the top of the list will be returned.
// Once an item is chosen, it will be moved to the bottom of the list.
func Choose(listfilePath string) (string, error) {

	var choice string

	list, loadErr := listfile.LoadListFromFile(listfilePath)
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
		rand.Seed(time.Now().UnixNano())

		// We want to make a random choice, but the items at the top of the list should have a better chance of being picked

		// First, we want to split the list in two
		thirdOfListLength := int(math.Floor(float64(len(list) / 3))) // We want to be sure we round down here
		priorityList := list[:thirdOfListLength]
		nonPriorityList := list[thirdOfListLength:]

		// We then choose a list to select from (3/4 chance it's the priority one)
		var chosenList []string
		if rand.Intn(4) == 0 {
			chosenList = nonPriorityList
		} else {
			chosenList = priorityList
		}

		// Now choose the item
		choice = chosenList[rand.Intn(len(chosenList))]
	}

	list = strlist.MoveToEnd(list, choice)
	saveErr := listfile.ReplaceListInFile(listfilePath, list)
	if saveErr != nil {
		return "", errors.New("Error while saving list: " + loadErr.Error())
	}

	return "Meditate on: " + choice, nil
}
