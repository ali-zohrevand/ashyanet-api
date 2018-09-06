package DB

import (
	"fmt"
	"testing"
)

func TestFindInArray(t *testing.T) {
	var arr []string
	arr = append(arr, "apple", "orange", "Good")

	i := FindInArray("apple", arr)
	if i != 0 {
		t.Error(i)
		t.Fail()
	}

}
func TestIsInArray(t *testing.T) {
	var arr []string
	arr = append(arr, "apple", "orange", "Good")

	f := IsInArray("apple", arr)
	if !f {
		t.Fail()
	}
}
func TestDeleteSliceByIndex(t *testing.T) {
	var arr []string
	arr = append(arr, "apple", "orange", "Good")

	a, err := DeleteSliceByIndex(1, arr)
	fmt.Println(a)
	if err != nil {
		t.Fail()
		t.Error(err)
	}
}
func TestDeleteSliceByObject(t *testing.T) {
	var arr []string
	arr = append(arr, "apple", "orange", "Good")

	a, err := DeleteSliceByObject("apple", arr)
	fmt.Println(a)
	if err != nil {
		t.Fail()
		t.Error(err)
	}
}
func TestDeleteRepetedCell(t *testing.T) {
	var arr []string
	arr = append(arr, "a", "a", "c", "c", "d", "a", "f")

	a, e := DeleteRepetedCell(arr)
	fmt.Println(a, e)
	if e != nil {
		t.Fail()
		t.Error(e)
	}
}
