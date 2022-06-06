package strlist

import "strings"

// Contains returns true if the given item exists in the given list. Returns false otherwise
// This is case-insensitive
func Contains(list []string, itemToFind string) bool {
	for _, item := range list {
		if strings.EqualFold(item, itemToFind) {
			return true
		}
	}

	return false
}

// RemoveFirstInstance retuns a copy of the given slice, with the first instance of the given item removed.
// If item isn't found, returns original list.
// This is case-insensitive
func RemoveFirstInstance(list []string, itemToRemove string) []string {

	matchFound := false
	for i := 0; i < len(list); i++ {
		if !matchFound && strings.EqualFold(list[i], itemToRemove) {
			matchFound = true
		}

		//If we've found it, start moving the next value into the current index
		if matchFound {
			if i < (len(list) - 1) {
				list[i] = list[i+1]
			}
		}
	}

	//If we've not found a match, return original array
	//Otherwise, the last item will either be the match, or will have been copied to the position before it (so ignore the last index)
	if !matchFound {
		return list
	}
	return list[:len(list)-1]
}

// MoveToEnd returns a copyof the given slice, with the first instance of the given item moved to the end of the slice.
// If item isn't found, returns original list.
// This is case-insensitive
func MoveToEnd(list []string, itemToMove string) []string {

	listOriginalLen := len(list)

	list = RemoveFirstInstance(list, itemToMove)
	if len(list) >= listOriginalLen {
		return list //If it wasn't found, return as is
	}

	list = append(list, itemToMove)
	return list
}

// FindBySubstring returns  slice, containing any string that contains the given substring.
// This is case-insensitive
func FindBySubstring(list []string, substringToFind string) []string {

	matches := []string{}

	for _, item := range list {
		if strings.Contains(strings.ToLower(item), strings.ToLower(substringToFind)) {
			matches = append(matches, item)
		}
	}

	return matches
}
