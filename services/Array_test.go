package services

import (
	"fmt"
	"testing"
)

func TestFindInArray(t *testing.T) {
	var arr []string
	arr = append(arr, "apple", "orange", "Good")
	s := make([]interface{}, len(arr))
	for i, v := range arr {
		s[i] = v
	}
	i := FindInArray("apple", s)
	if i != 0 {
		t.Error(i)
		t.Fail()
	}

}
func TestIsInArray(t *testing.T) {
	var arr []string
	arr = append(arr, "apple", "orange", "Good")
	s := make([]interface{}, len(arr))
	for i, v := range arr {
		s[i] = v
	}
	f := IsInArray("apple", s)
	if !f {
		t.Fail()
	}
}
func TestDeleteSliceByIndex(t *testing.T) {
	var arr []string
	arr = append(arr, "apple", "orange", "Good")
	s := make([]interface{}, len(arr))
	for i, v := range arr {
		s[i] = v
	}
	a, err := DeleteSliceByIndex(1, s)
	fmt.Println(a)
	if err != nil {
		t.Fail()
		t.Error(err)
	}
}
func TestDeleteSliceByObject(t *testing.T) {
	var arr []string
	arr = append(arr, "apple", "orange", "Good")
	s := make([]interface{}, len(arr))
	for i, v := range arr {
		s[i] = v
	}
	a, err := DeleteSliceByObject("apple", s)
	fmt.Println(a)
	if err != nil {
		t.Fail()
		t.Error(err)
	}
}
