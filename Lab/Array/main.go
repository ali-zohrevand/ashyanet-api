package main

import "fmt"

func main() {
	var a []string
	a = append(a, "asdsd", "sdasd", "asdasd", "asdsad")
	fmt.Println(a)
	for e, s := range a {
		fmt.Println(e, " ", s)
	}

}
