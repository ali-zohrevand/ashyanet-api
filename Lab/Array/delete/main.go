package main

import "fmt"

func main() {
	mySlice := make([]int, 0, 5) //make (T,length,capacity)
	fmt.Println("...............................")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("...............................")
	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len: ", len(mySlice), " Capacity: ", cap(mySlice), " Value: ", mySlice[i])
	}

	slice := []int{1, 6, 666666, 55, 44}
	mySlice = append(mySlice, slice...) // add slice to anoeher slice
	fmt.Println("...............................")
	fmt.Println(mySlice)
	//we want delete i'th element
	fmt.Println("...............................")
	fmt.Println(slice)
	slice = append(slice[:2], slice[3:]...)
	fmt.Println("...............................")
	fmt.Println(slice)
	fmt.Println("...............................")
	slice, _ = DeleteSlice(1, slice)
	fmt.Println(slice)

}
func DeleteSlice(i int, slice []int) ([]int, bool) {
	tempLenght := len(slice)
	if i > len(slice) || i < 0 {
		return slice, false
	}

	slice = append(slice[:i], slice[i+1:]...)
	tempLenght--
	if len(slice) == tempLenght {
		return slice, true

	} else {
		return slice, false
	}
}
