package listfile

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// Append to or replace list in file. isAppending parameter determines whether the file is added to, or whether the contents are replaced
func writeTofile(list []string, file *os.File, isAppending bool) error {
	w := bufio.NewWriter(file)

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	//If file is empty, and we're appending, there is still technically the first line to think about. Therefore if file is empty, don't append
	if fileInfo.Size() == 0 {
		isAppending = false
	}

	if isAppending {
		fmt.Fprint(w, "\n")
	}

	for i, item := range list {

		fmt.Fprint(w, item)

		if i < (len(list) - 1) {
			fmt.Fprint(w, "\n")
		}
	}

	return w.Flush()
}

// LoadListFromFile will load the contents from a file, as a slice of strings
func LoadListFromFile(path string) ([]string, error) {

	file, err := os.OpenFile(path, os.O_CREATE, 0644)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	var list []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}

	return list, scanner.Err()
}

// ReplaceListInFile will replace the contents in a file, as a slice of strings
func ReplaceListInFile(path string, list []string) error {

	file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	return writeTofile(list, file, false)
}

// AppendToListInFile will append to the contents of a file, as a slice of strings
func AppendToListInFile(path string, newItems []string) error {

	fileIsNotNew := true
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		fileIsNotNew = false
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	return writeTofile(newItems, file, fileIsNotNew)

}
