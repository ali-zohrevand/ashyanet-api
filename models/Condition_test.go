package models

import (
	"fmt"
	"testing"
)

func TestCondition_GetAttr(t *testing.T) {

}
func TestCondition_Happened(t *testing.T) {
	var EqualeJsonString Condition

	EqualeJsonString.InData = `{"some":"sc"}`
	EqualeJsonString.JsonAttributeName = "some"
	EqualeJsonString.ConditionType = EqualeTo
	EqualeJsonString.GetAttr("sc")
	var EqualJsonInt Condition
	EqualJsonInt.JsonAttributeName = "some"
	EqualJsonInt.ConditionType = EqualeTo
	EqualJsonInt.InData = `{"some":1}`
	EqualJsonInt.GetAttr(1)

	var StringTestEquale Condition
	StringTestEquale.InData = "a"
	StringTestEquale.ConditionType = EqualeTo
	a := "a"
	StringTestEquale.GetAttr(a)
	var StringTestNotEquale Condition
	StringTestNotEquale.InData = "a"
	StringTestNotEquale.ConditionType = EqualeTo
	b := "b"
	StringTestNotEquale.GetAttr(b)
	var testIntCondtionGraterTHan Condition
	testIntCondtionGraterTHan.InData = 9
	testIntCondtionGraterTHan.ConditionType = GraterThan
	testIntCondtionGraterTHan.GetAttr(5)
	var testIntCondtionLowerTHan Condition
	testIntCondtionLowerTHan.InData = 1
	testIntCondtionLowerTHan.ConditionType = LowerThan
	testIntCondtionLowerTHan.GetAttr(5)
	var testIntCondtionbetween Condition
	testIntCondtionbetween.InData = 2
	testIntCondtionbetween.ConditionType = Between
	testIntCondtionbetween.GetAttr(1)
	testIntCondtionbetween.GetAttr(9)
	var testIntCondtionbetweenFalse Condition
	testIntCondtionbetweenFalse.InData = 10
	testIntCondtionbetweenFalse.ConditionType = Between
	testIntCondtionbetweenFalse.GetAttr(1)
	testIntCondtionbetweenFalse.GetAttr(9)

	var tests = []struct {
		input Condition
		Is    bool
		err   error
	}{
		{EqualJsonInt, true, nil},
		{testIntCondtionbetween, true, nil},
		{testIntCondtionbetweenFalse, false, nil},
		{EqualeJsonString, true, nil},
		{StringTestEquale, true, nil},
		{StringTestNotEquale, false, nil},
		{testIntCondtionGraterTHan, true, nil},
		{testIntCondtionLowerTHan, true, nil},
	}
	for i, test := range tests {
		ok, err := test.input.Happened()
		if ok != test.Is || err != test.err && test.input.InData != "" {
			fmt.Println(i+1, "\n =========")
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.Is, test.err)
			t.Fail()
		}
	}
}
