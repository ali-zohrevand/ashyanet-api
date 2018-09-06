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
	for i, v := range array {
		count := NumberInArray(v, out)
		if count > 1 {

			for j, a := range out {
				if a == v && i != j {
					out, err = DeleteSliceByIndex(j, out)
					count = NumberInArray(v, out)
					if err != nil {
						return
					}
				}
			}
		}
	}
	return
}
