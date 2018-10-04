package models

import (
	"fmt"
	"testing"
)

func TestCondition_GetAttr(t *testing.T) {

}
func TestCondition_Happened(t *testing.T) {
	var EqualeJsonString Condition
	EqualeJsonString.JsonAttributeName = "some"
	EqualeJsonString.ConditionType = EqualeTo
	EqualeJsonString.GetAttr("sc")
	var EqualJsonInt Condition
	EqualJsonInt.JsonAttributeName = "some"
	EqualJsonInt.ConditionType = EqualeTo
	EqualJsonInt.GetAttr(1)

	var StringTestEquale Condition
	StringTestEquale.ConditionType = EqualeTo
	a := "a"
	StringTestEquale.GetAttr(a)
	var StringTestNotEquale Condition
	StringTestNotEquale.ConditionType = EqualeTo
	b := "b"
	StringTestNotEquale.GetAttr(b)
	var testIntCondtionGraterTHan Condition
	testIntCondtionGraterTHan.ConditionType = GraterThan
	testIntCondtionGraterTHan.GetAttr(5)
	var testIntCondtionLowerTHan Condition
	testIntCondtionLowerTHan.ConditionType = LowerThan
	testIntCondtionLowerTHan.GetAttr(5)
	var testIntCondtionbetween Condition
	testIntCondtionbetween.ConditionType = Between
	testIntCondtionbetween.GetAttr(1)
	testIntCondtionbetween.GetAttr(9)
	var testIntCondtionbetweenFalse Condition
	testIntCondtionbetweenFalse.ConditionType = Between
	testIntCondtionbetweenFalse.GetAttr(1)
	testIntCondtionbetweenFalse.GetAttr(9)

	var tests = []struct {
		c     Condition
		input interface{}
		Is    bool
		err   error
	}{
		{EqualJsonInt, `{"some":1}`, true, nil},
		{testIntCondtionbetween, 7, true, nil},
		{testIntCondtionbetweenFalse, 0, false, nil},
		{EqualeJsonString, `{"some":"sc"}`, true, nil},
		{StringTestEquale, "a", true, nil},
		{StringTestNotEquale, "s", false, nil},
		{testIntCondtionGraterTHan, 9, true, nil},
		{testIntCondtionLowerTHan, 1, true, nil},
	}
	for i, test := range tests {
		ok, err := test.c.Happened(test.input)
		if ok != test.Is || err != test.err {
			fmt.Println(i+1, "\n =========")
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.c, test.Is, test.err)
			t.Fail()
		}
	}
}
