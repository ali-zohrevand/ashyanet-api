package DB

import "errors"

func IsInArray(Target string, Array []string) (Found bool) {

	i := FindInArray(Target, Array)
	if i == -1 {
		return false
	} else {
		return true
	}
	return
}
func DeleteSliceByIndex(i int, slice []string) (output []string, err error) {
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
func DeleteSliceByObject(Target string, slice []string) (output []string, err error) {
	i := FindInArray(Target, slice)
	if i == -1 {
		return nil, errors.New("OBJECT NOT FOUND")
	} else {
		output, err = DeleteSliceByIndex(i, slice)
	}
	return

}
func FindInArray(Target string, Array []string) (Index int) {
	for i, a := range Array {
		if a == Target {
			return i
		}
	}
	return -1
}
func FindAllInArra(Target string, Array []string) (Index []int) {
	for i, a := range Array {
		if a == Target {
			Index = append(Index, i)
		}
	}
	return Index
}
func NumberInArray(Target string, Array []string) (count int) {
	count = 0
	for _, a := range Array {
		if a == Target {
			count++
		}
	}
	return count
}
func DeleteRepetedCell(array []string) (out []string, err error) {
	for _, b := range array {
		out = append(out, b)
	}
	for _, a := range out {
		for {
			Count := FindAllInArra(a, array)
			if len(Count) > 1 {
				array, _ = DeleteSliceByIndex(Count[1], array)
			} else {
				break
			}

		}
	}
	return array, nil
}
func AddToArrayUnique(object string, Array []string) []string {
	All := FindAllInArra(object, Array)
	if len(All) > 1 {
		return Array
	}
	Array = append(Array, object)
	return Array
}
