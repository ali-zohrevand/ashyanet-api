package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	//add:="/Users/alizohrevand/homebrew/Cellar/mosquitto/1.4.14_2/etc/mosquitto"
	//passAdd:=path.Join(add,"pass")
	//_,err:=createFile(passAdd)
	cmd := "mosquitto_passwd -b pass A pass"
	//cmd:="ls"
	createFile("pass")
	cdcomand := exec.Command("sh", "-c", cmd)
	cmdOutput := &bytes.Buffer{}
	cdcomand.Stdout = cmdOutput
	//cdcomand.Path=add
	err := cdcomand.Run()
	if err != nil {
		fmt.Printf("%s", err)
	}
	out := cmdOutput.Bytes()
	fmt.Println(BytesToString(out))

	return
}
func BytesToString(data []byte) string {
	return string(data[:])
}
func createFile(Addr string) (file *os.File, err error) {
	file, err = os.Open(Addr)
	if err != nil {
		file, err = os.Create(Addr)
	}
	return
}
