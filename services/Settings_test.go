package services

import "testing"

func TestCreateSettingsFile(t *testing.T) {
	err := CreateSettingsFile()
	if err != nil {
		t.Fail()
		t.Error(err)
	}
}
