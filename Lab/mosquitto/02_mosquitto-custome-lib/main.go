package main

import (
	"SimpleAPIBasePlatform/Doc/mosquitto/02_mosquitto-custome-lib/msquitto"
	"fmt"
)

/*	linux conf Path: /etc/mosquitto/conf.d/default.conf
	restart command: systemctl restart mosquitto
	add user: mosquitto_passwd
//..............
	path:="/Users/alizohrevand/homebrew/Cellar/mosquitto/1.4.14_2/etc/mosquitto/mosquitto.conf"
	restart_command:="brew services start mosquitto"
*/
func main() {
	path := "/etc/mosquitto/conf.d/default.conf"
	restart_command := "sudo systemctl restart mosquitto"
	add_user := "mosquitto_passwd"
	mos := msquitto.NewMosquitto(path, restart_command, add_user)
	err := mos.Deny_anonymous()
	fmt.Println("err: ", err)
	err = mos.AddUser("ali", "ali")
	fmt.Println("err: ", err)
	_, err = mos.RestartMosquitto()
	fmt.Println("err: ", err)

}
