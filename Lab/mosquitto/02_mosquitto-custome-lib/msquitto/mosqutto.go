package msquitto

import (
	"bytes"
	"os"
	"os/exec"
	"strings"

	"fmt"
	"io/ioutil"
	"path"
	"runtime"
)

type mosquitto struct {
	confAdd        string
	restartCommand string
	addUserCommand string
}

func NewMosquitto(confDir string, re string, add string) *mosquitto {
	return &mosquitto{confAdd: confDir, restartCommand: re, addUserCommand: add}
}
func (mos *mosquitto) Deny_anonymous() (err error) {
	Disable := "\nallow_anonymous false"
	err = putRulesInMosquittoConf(mos.confAdd, Disable)
	return err
}
func (mos *mosquitto) AddUser(username string, password string) (err error) {

	_, err = createFile("pass")
	if err != nil {
		return
	}
	_, filename, _, _ := runtime.Caller(1)
	PassAddr := path.Join(path.Dir(filename), "pass")
	stringToAdd := "\npassword_file " + PassAddr
	putRulesInMosquittoConf(mos.confAdd, stringToAdd)
	_, err = exe(mos.addUserCommand + " -b pass " + username + " " + password)
	if err != nil {
		return
	}
	return
}
func createFile(Addr string) (file *os.File, err error) {
	file, err = os.Open(Addr)
	if err != nil {
		file, err = os.Create(Addr)
	}
	return
}
func putRulesInMosquittoConf(add string, rule string) (err error) {

	b, err := ioutil.ReadFile(add) // just pass the file name
	fmt.Println(err)
	file, err := os.OpenFile(add, os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
		}
	}()
	ConfInString := string(b)
	if strings.Contains(ConfInString, rule) {
		return
	} else {
		file.WriteString(rule)
	}

	return
}
func (mos *mosquitto) RestartMosquitto() (OutPut string, err error) {

	return exe(mos.restartCommand)
}

func exe(cmd string) (OutPut string, err error) {
	cdcomand := exec.Command("sh", "-c", cmd)
	cmdOutput := &bytes.Buffer{}
	cdcomand.Stdout = cmdOutput
	err = cdcomand.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
	OutPut = bytesToString(cmdOutput.Bytes())
	return
}
func bytesToString(data []byte) string {
	return string(data[:])
}
