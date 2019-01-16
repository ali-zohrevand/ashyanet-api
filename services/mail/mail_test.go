package simple_mail

import (
	"fmt"
	"github.com/ali-zohrevand/ashyanet-api/settings/Words"

	"testing"
)

func TestMAIL(t *testing.T) {

	mailTosend := NewMail("smtp.gmail.com", "465", "ashyanet@gmail.com", "mahdi1369QWE", []string{"alihooshyar1990@gmail.com"}, nil, nil, "new message", Words.VerifyMail, false)
	success, err := mailTosend.SendMail()
	fmt.Println(success, err)
	if err != nil {
		t.Fail()
		t.Error(err)
	}

}
