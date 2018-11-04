package services

import (
	"errors"
)

func IsInArray(Target interface{}, Array []interface{}) (Found bool) {

	i := FindInArray(Target, Array)
	if i == -1 {
		return false
	} else {
		return true
	}
	return
}
func DeleteSliceByIndex(i int, slice []interface{}) (output []interface{}, err error) {
	tempLenght := len(slice)
	if i > len(slice) || i < 0 {
		err = errors.New("OUT OF RANGE")
		output = slice
		return
	}

	slice = append(slice[:i], slice[i+1:]...)
	tempLenght--
	if len(slice) == tempLenght {
		return slice, nil

	} else {
		return slice, errors.New("OUT OF RANGE")
	}
}

func DeleteSliceByObject(Target interface{}, slice []interface{}) (output []interface{}, err error) {
	i := FindInArray(Target, slice)
	if i == -1 {
		return nil, errors.New("OBJECT NOT FOUND")
	} else {
		output, err = DeleteSliceByIndex(i, slice)
	}
	return

}
func FindInArray(Target interface{}, Array []interface{}) (Index int) {
	for i, a := range Array {
		if a == Target {
			return i
		}
	}
	return -1
}
func NumberInArray(Target interface{}, Array []interface{}) (count int) {
	count = 0
	for _, a := range Array {
		if a == Target {
			count++
		}
	}
	return count
}
func DeleteRepetedCell(array []interface{}) (out []interface{}, err error) {
	for _, b := range array {
		out = append(out, b)
	}
	for i, v := range array {
		count := NumberInArray(v, out)
		if count > 1 {
			for j, a := range out {
				if a == v && i != j {
					out, err = DeleteSliceByIndex(j, out)
					if err != nil {
						return
					}
				}
			}
		}
	}
	return
}
