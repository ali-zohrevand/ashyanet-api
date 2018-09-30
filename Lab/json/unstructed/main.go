package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	myJsonString := `{"some":"json"}`

	// `&myStoredVariable` is the address of the variable we want to store our
	// parsed data in
	var myStoredVariable interface{}
	err := json.Unmarshal([]byte(myJsonString), &myStoredVariable)
	valid := json.Valid([]byte(myJsonString))
	fmt.Println(myStoredVariable)
	fmt.Println(valid)
	fmt.Println(err)
}
