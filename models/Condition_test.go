package models

import (
	"fmt"
	"testing"
)

func TestCondition_Happened(t *testing.T) {
	var c Condition
	var d []string
	d = append(d, "s", "sd")
	c.InData = d
	s, e := c.Happened()
	fmt.Println(s, e)

}
