package main

import (
	"fmt"
	"github.com/asaskevich/govalidator"
)

func main() {
	type Post struct {
		Title    string `valid:"alphanum,required~First name is blank,runelength(5|10)~length not good"`
		Message  string `valid:"duck,ascii"`
		AuthorIP string `valid:"ipv4~Ip Error"`
		Date     string `valid:"-"`
	}
	post := &Post{
		Title:    "s",
		Message:  "s",
		AuthorIP: "123.2354.3",
	}

	// Add your own struct validation tags
	govalidator.TagMap["duck"] = govalidator.Validator(func(str string) bool {
		return str == "duck"
	})
	var temp = "asd"
	result, err := govalidator.ValidateStruct(post)
	if err != nil {
		errors := err.(govalidator.Errors).Errors()
		for _, e := range errors {
			fmt.Println("", e.Error())
		}
	}
	println(result)
	test := govalidator.RuneLength(temp, "3", "7")
	fmt.Println("test ", test)
	BadCharList := []string{",", "/", "<", ">", "$", "'", "!", ")", "(", "&", "%", "~", "=", "+", "-", "?"}
	for badChar, a := range BadCharList {
		fmt.Println(badChar, " ", a)
	}
}
