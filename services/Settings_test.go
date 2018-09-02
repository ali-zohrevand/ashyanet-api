package services

import "testing"

func TestCreateSettingsFile(t *testing.T) {
	err := CreateSettingsFile()
	if err != nil {
		t.Fail()
		t.Error(err)
	}
}
func TestSaveSteetinInDB(t *testing.T) {
	err := SaveSetingsInDB()
	if err != nil {
		t.Fail()
		t.Error(err)
	}
}
