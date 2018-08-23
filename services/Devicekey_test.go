package services

import (
	"fmt"
	"testing"
)

func TestGetValidKey(t *testing.T) {
	code, message := CreatValidKey()
	if code != 200 {
		t.Error(BytesToString(message))
	} else {
		fmt.Println(BytesToString(message))
	}
}
