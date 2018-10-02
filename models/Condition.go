package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"gitlab.com/hooshyar/ChiChiNi-API/settings/Words"
	"reflect"
)

type Condition struct {
	InData            interface{}
	JsonAttributeName string
	CommandFunction   Command
	ConditionType     ConditionType
}

type ConditionType int

const (
	LowerThan  ConditionType = 0
	GraterThan ConditionType = 1
	Between    ConditionType = 2
	EqualeTo   ConditionType = 3
)

func (c *Condition) Happened(Boundries ...interface{}) (Ok bool, err error) {
	typeOdData := reflect.TypeOf(c.InData).String()
	if IsJson(c.InData) {
		typeOdData = "json"
		var s error
		c.InData, typeOdData, s = GetDataFromJsom(c.InData, c.JsonAttributeName)
		if s != nil && typeOdData == "" {
			return false, errors.New(Words.InvalidaData)
		}
	}
	IsDataTypeOK, dataTypeDetected := ConditionIsDataTypeOK(typeOdData, c.InData)
	var BType string
	if len(Boundries) != 0 {
		BType = reflect.TypeOf(Boundries[0]).String()
	}
	if BType != typeOdData {
		return false, errors.New(Words.InvalidaData)
	}
	if !ConditionIsBoundriesLenghtOk(c.ConditionType, c.InData, len(Boundries)) || c.InData == nil || !IsDataTypeOK {
		return false, errors.New(Words.InvalidaData)
	}
	fmt.Println(dataTypeDetected)
	switch c.ConditionType {
	case GraterThan:
		if dataTypeDetected == "string" {
			return false, errors.New(Words.InvalidaData)
		}
		Check := Boundries[0]
		if IsGraterThan(c.InData, Check) {
			return true, nil
		} else {
			return false, nil
		}

	case Between:
		if dataTypeDetected == "string" {
			return false, errors.New(Words.InvalidaData)
		}
		a := Boundries[0]
		b := Boundries[1]
		if IsBetween(c.InData, a, b) {
			return true, nil
		} else {
			return false, nil
		}
	case EqualeTo:
		Check := Boundries[0]
		if IsEquale(c.InData, Check) {
			return true, nil
		} else {
			return false, nil
		}
	case LowerThan:
		if dataTypeDetected == "string" {
			return false, errors.New(Words.InvalidaData)
		}
		Check := Boundries[0]
		if IsLowerThan(c.InData, Check) {
			return true, nil
		} else {
			return false, nil
		}
	}
	return
}
func GetDataFromJsom(JsonString interface{}, Key string) (Resault interface{}, dataType string, err error) {
	var JsonMap map[string]interface{}
	val := fmt.Sprintf("%v", JsonString)
	errJson := json.Unmarshal([]byte(val), &JsonMap)
	if errJson != nil {
		return "", "", errors.New(Words.InvalidaData)
	}
	dataInJson := JsonMap[Key]
	if dataInJson == nil {
		return "", "", errors.New(Words.InvalidaData)
	}
	value := fmt.Sprintf("%v", dataInJson)
	x, err := strconv.Atoi(value)
	if err != nil {
		Resault = dataInJson
		dataType = reflect.TypeOf(dataInJson).String()
	} else {
		Resault = x
		dataType = reflect.TypeOf(x).String()
	}
	return
}
func IsString(in interface{}) bool {
	if reflect.TypeOf(in).String() == "string" {
		return true
	}
	return false
}
func IsEquale(a interface{}, b interface{}) bool {
	TypeOfA := reflect.TypeOf(a).String()
	TypeOfB := reflect.TypeOf(b).String()

	if TypeOfA != TypeOfB {
		return false
	}
	if a == b {
		return true
	}
	return false
}
func IsJson(in interface{}) bool {

	val := fmt.Sprintf("%v", in)
	valid := json.Valid([]byte(val))
	if valid && !IsInt(in) && !IsBool(in) {
		return true
	}
	return false
}
func IsLowerThan(in interface{}, lenght interface{}) bool {
	if !IsInt(in) || !IsInt(lenght) {
		return false
	}
	a := fmt.Sprintf("%v", in)
	b := fmt.Sprintf("%v", lenght)

	var x, y int
	var err error
	x, err = strconv.Atoi(a)
	if err != nil {
		return false
	}
	y, err = strconv.Atoi(b)
	if err != nil {
		return false
	}
	if x < y {
		return true
	}
	return false
}

func IsEqule(in interface{}, lenght interface{}) bool {
	a := fmt.Sprintf("%v", in)
	b := fmt.Sprintf("%v", lenght)
	var x, y int
	var err error
	x, err = strconv.Atoi(a)
	if err != nil {
		return false
	}
	y, err = strconv.Atoi(b)
	if err != nil {
		return false
	}
	if x == y {
		return true
	}
	if a == b {
		return true
	}
	return false
}

func IsBetween(input interface{}, low interface{}, up interface{}) bool {
	if !IsInt(input) || !IsInt(low) || !IsInt(up) {
		return false
	}
	i := fmt.Sprintf("%v", input)
	x := fmt.Sprintf("%v", low)
	y := fmt.Sprintf("%v", up)
	var err error
	var in, a, b int
	in, err = strconv.Atoi(i)
	if err != nil {
		return false
	}
	a, err = strconv.Atoi(x)
	if err != nil {
		return false
	}
	b, err = strconv.Atoi(y)
	if err != nil {
		return false
	}
	if in > a && in < b {
		return true
	}
	return false
}
func IsGraterThan(in interface{}, lenght interface{}) bool {
	if !IsInt(in) || !IsInt(lenght) {
		return false
	}
	a := fmt.Sprintf("%v", in)
	b := fmt.Sprintf("%v", lenght)
	var x, y int
	var err error
	x, err = strconv.Atoi(a)
	if err != nil {
		return false
	}
	y, err = strconv.Atoi(b)
	if err != nil {
		return false
	}
	if x > y {
		return true
	}
	return false
}
func IsInt(in interface{}) bool {
	if reflect.TypeOf(in).String() == "int" {
		return true
	}
	return false
}
func IsBool(in interface{}) bool {
	if reflect.TypeOf(in).String() == "bool" {
		return true
	}
	return false
}
func ConditionIsDataTypeOK(dataType string, InData interface{}) (ok bool, typeDected string) {
	switch dataType {
	case "int":
		return true, "int"
	case "string":
		return true, "string"
	case "bool":
		return false, "bool"

	case "json":
		val := fmt.Sprintf("%v", InData)
		valid := json.Valid([]byte(val))
		if valid {
			var JsonMap interface{}
			errJson := json.Unmarshal([]byte(val), &JsonMap)
			if errJson != nil {
				return false, ""
			}
			return true, "json"
		}
		return false, ""

	default:
		return false, "UNKOWN"
	}
	return false, ""
}
func ConditionIsBoundriesLenghtOk(conditionType ConditionType, InDatea interface{}, BoundiresLntght int) (ok bool) {
	switch conditionType {
	case LowerThan:
		if BoundiresLntght != 1 {
			return false
		}
	case GraterThan:
		if BoundiresLntght != 1 {
			return false
		}
	case EqualeTo:
		if BoundiresLntght != 1 {
			return false
		}
	case Between:
		if BoundiresLntght != 2 {
			return false
		}
	default:
		return false
	}

	return true
}
