package services

import "errors"

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
