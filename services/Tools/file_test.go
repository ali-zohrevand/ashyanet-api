package Tools

import (
	"fmt"
	"testing"
)

func TestCreateFile(t *testing.T) {
	path := "test.txt"
	err := CreateFile(path)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
func TestWriteFile(t *testing.T) {
	path := "test.txt"
	err := WriteFile(path, "Salam")
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	err = DeleteFile(path)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
func TestReadFile(t *testing.T) {
	path := "test.txt"
	err := CreateFile(path)
	err = WriteFile(path, "Salam")
	m, err := ReadFile(path)
	if m != "Salam" {
		fmt.Println(m)
		t.Error(err)
		t.Fail()
	}
	err = DeleteFile(path)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
func TestDeleteFile(t *testing.T) {
	path := "test.txt"
	err := CreateFile(path)
	err = DeleteFile(path)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
