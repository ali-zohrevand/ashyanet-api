package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func main() {
	conDir := "/Users/alizohrevand/homebrew/Cellar/mosquitto/1.4.14_2"
	confFile := "test2"
	confAddr := path.Join(conDir, confFile)
	d1 := []byte("hello\ngo\n")
	/*	err := ioutil.WriteFile(confAddr, d1, 0644)
		fmt.Println(err)*/
	fo, err := os.Open(confAddr)
	defer fo.Close()
	if err != nil {
		fmt.Println(err)
	}
	fo.Write(d1)
	buffer := make([]byte, 10)
	fo.Read(buffer)
	fmt.Println(string(buffer))
	temp := string(buffer)
	if strings.Contains(temp, "go") {
		temp = strings.Replace(temp, "\ngo", "bye", -1)

	} else {
		temp = temp + "\n" + "bye"
	}
	b := []byte(temp)

	fmt.Println("last ", err)
}
func createFile(Addr string) (file *os.File, err error) {
	file, err = os.Open(Addr)
	if err != nil {
		file, err = os.Create(Addr)

	}
	return
}
