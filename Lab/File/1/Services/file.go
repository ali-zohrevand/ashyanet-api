package Services

import (
	"fmt"
	"io"
	"os"
)

func CreateFile(FilePath string) {
	// check if file exists
	var _, err = os.Stat(FilePath)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(FilePath)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("File Created Successfully", FilePath)
}

func WriteFile(FilePath string, input string) {
	// Open file using READ & WRITE permission.
	var file, err = os.OpenFile(FilePath, os.O_RDWR|os.O_APPEND, 0660)
	if isError(err) {
		return
	}
	defer file.Close()

	// Write some text line-by-line to file.
	_, err = file.WriteString(input)
	if isError(err) {
		return
	}

	// Save file changes.
	err = file.Sync()
	if isError(err) {
		return
	}

}

func readFile(FilePath string) (Content string, err error) {
	// Open file for reading.
	file, err := os.OpenFile(FilePath, os.O_RDWR, 0660)
	if isError(err) {
		return
	}
	defer file.Close()

	// Read file, line by line
	var text = make([]byte, 1024)
	for {
		_, err = file.Read(text)

		// Break if finally arrived at end of file
		if err == io.EOF {
			break
		}

		// Break if error occured
		if err != nil && err != io.EOF {
			isError(err)
			break
		}
	}

	Content = string(text)
	return
}

func deleteFile(FilePath string) (err error) {
	// delete file
	err = os.Remove(FilePath)
	if isError(err) {
		return
	}
	return

}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
