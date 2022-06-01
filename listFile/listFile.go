package listFile

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func writeTofile(list []string, file *os.File, isAppending bool) error {
	w := bufio.NewWriter(file)

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

func ReplaceListInFile(path string, list []string) error {

	file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	return writeTofile(list, file, false)
}

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
