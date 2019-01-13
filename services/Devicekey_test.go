package services

import (
	"fmt"
	"gitlab.com/hooshyar/ChiChiNi-API/services/Tools"
	"testing"
)

func TestGetValidKey(t *testing.T) {
	code, message := CreatValidKey()
	if code != 200 {
		t.Error(Tools.BytesToString(message))
	} else {
		fmt.Println(Tools.BytesToString(message))
	}
}
