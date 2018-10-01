package models

import (
	"fmt"
	"testing"
)

func TestCondition_Happened(t *testing.T) {
	var c Condition

	c.InData = `{"some":4}`
	c.JsonAttributeName = "some"
	c.ConditionType = GraterThan
	s, e := c.Happened(1)
	fmt.Println(s, e)

}
