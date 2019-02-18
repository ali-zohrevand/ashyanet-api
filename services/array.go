package services

import "errors"

func DeleteRepetedCell(array []string) (out []string) {
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
	return array
}
func FindAllInArra(Target string, Array []string) (Index []int) {
	for i, a := range Array {
		if a == Target {
			Index = append(Index, i)
		}
	}
	return Index
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
