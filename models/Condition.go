package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"gitlab.com/hooshyar/ChiChiNi-API/settings/Words"
	"reflect"
)

type Condition struct {
	InData          interface{}
	CommandFunction Command
	ConditionType   ConditionType
}

type ConditionType int

const (
	LowerThan  ConditionType = 0
	GraterThan ConditionType = 1
	Between    ConditionType = 2
	EqualeTo   ConditionType = 3
)

func (c *Condition) Happened(Boundries ...interface{}) (Ok bool, err error) {

	if !isBoundriesOk(c.ConditionType, Boundries) {
		return false, errors.New(Words.InvalidaData)
	}
	if c.InData == nil {
		return false, errors.New(Words.InvalidaData)
	}
	typeOdData := reflect.TypeOf(c.InData)
	switch typeOdData.String() {
	case "int":

	case "string":
	case "bool":
		return false, errors.New(Words.InvalidaData)

	default:
		str := fmt.Sprintf("%v", c.InData)
		valid := json.Valid([]byte(str))
		if valid {
			var JsonMap interface{}
			errJson := json.Unmarshal([]byte(str), &JsonMap)
			if errJson != nil {
				return false, errors.New(Words.InvalidaData)
			}
			fmt.Println("data is json")

		}
		return false, errors.New(Words.InvalidaData)

	}

	return
}
func isBoundriesOk(conditionType ConditionType, Boundires ...interface{}) (ok bool) {
	switch conditionType {
	case 1:
		if len(Boundires) != 1 {
			return false
		}
	case 2:
		if len(Boundires) != 2 {
			return false
		}
	case 3:
		if len(Boundires) != 1 {
			return false
		}
	case 0:
		if len(Boundires) != 1 {
			return false
		}
	default:
		return
	}
	return
}
