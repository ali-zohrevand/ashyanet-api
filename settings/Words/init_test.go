package Words

import "testing"

func TestInitDbWOrds(t *testing.T) {
	_, err := InitDbWords()
	if err != nil {
		t.Fail()
		t.Error(err)
	}
}
func TestInitKeyWords(t *testing.T) {
	_, err := InitValidationWords()
	if err != nil {
		t.Fail()
		t.Error(err)
	}
}
