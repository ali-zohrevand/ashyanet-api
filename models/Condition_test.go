package models

import (
	"testing"
)

func TestCondition_Happened(t *testing.T) {
	var EqualeJsonString Condition

	EqualeJsonString.InData = `{"some":"sc"}`
	EqualeJsonString.JsonAttributeName = "some"
	EqualeJsonString.ConditionType = EqualeTo
	var EqualJsonInt Condition
	EqualeJsonString.JsonAttributeName = "some"
	EqualeJsonString.ConditionType = EqualeTo
	EqualJsonInt.InData = `{"some":1}`
	var tests = []struct {
		input     Condition
		Boundries interface{}
		Is        bool
		err       error
	}{
		{EqualeJsonString, "sc", true, nil},
		{EqualJsonInt, 1, true, nil},
	}
	for _, test := range tests {
		ok, err := EqualeJsonString.Happened(test.Boundries)
		if ok != test.Is && err != test.err {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.Is, test.err)
			t.Fail()
		}
	}
}
