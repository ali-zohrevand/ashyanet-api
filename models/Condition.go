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
	JsonAttributeName string        `json:"json_attribute_name" bson:"json_attribute_name"`
	ConditionType     ConditionType `json:"condition_type" bson:"condition_type"`
	Attr              []interface{} `json:"attr" bson:"attr"`
}

type ConditionType int

const (
	LowerThan  ConditionType = 0
	GraterThan ConditionType = 1
	Between    ConditionType = 2
	EqualeTo   ConditionType = 3
)

func (c *Condition) GetAttr(in interface{}) {

	c.Attr = append(c.Attr, in)

}
func (c *Condition) IsValid() (Is bool) {
	if len(c.Attr) > 2 || len(c.Attr) < 1 {
		return false
	}
	if c.ConditionType < 0 || c.ConditionType > 4 {
		return false
	}
	return true
}
func (c *Condition) Happened(Input string) (Ok bool, err error) {
	var Boundries []interface{}
	for _, v := range c.Attr {
		Boundries = append(Boundries, v)
	}
	var typeOdData string
	if IsJson(Input) {
		typeOdData = "json"
		var s error
		Input, typeOdData, s = GetDataFromJsom(Input, c.JsonAttributeName)
		if s != nil && typeOdData == "" {
			return false, errors.New(Words.InvalidaData)
		}
	}
	IsDataTypeOK, dataTypeDetected := ConditionIsDataTypeOK(typeOdData, Input)
	var BType string
	if len(Boundries) != 0 {
		BType = reflect.TypeOf(Boundries[0]).String()
	}
	if BType != typeOdData {
		return false, errors.New(Words.InvalidaData)
	}
	if !ConditionIsBoundriesLenghtOk(c.ConditionType, Input, len(Boundries)) || Input == nil || !IsDataTypeOK {
		return false, errors.New(Words.InvalidaData)
	}
	switch c.ConditionType {
	case GraterThan:
		if dataTypeDetected == "string" {
			return false, errors.New(Words.InvalidaData)
		}
		Check := Boundries[0]
		if IsGraterThan(Input, Check) {
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
		if IsBetween(Input, a, b) {
			return true, nil
		} else {
			return false, nil
		}
	case EqualeTo:
		Check := Boundries[0]
		if IsEquale(Input, Check) {
			return true, nil
		} else {
			return false, nil
		}
	case LowerThan:
		if dataTypeDetected == "string" {
			return false, errors.New(Words.InvalidaData)
		}
		Check := Boundries[0]
		if IsLowerThan(Input, Check) {
			return true, nil
		} else {
			return false, nil
		}
	}
	return
}
func (c *Condition)whatIsDataType(Input string)(out interface{},dataType string)  {
	intValue,isInt:=IsInt(Input)
	if isInt{
		return intValue,"int"
	}
	jsonValue,jasonType,IsJson:=IsJson()

	return
}
func GetDataFromJsom(JsonString string, Key string) (Resault interface{}, dataType string, err error) {
	var JsonMap map[string]interface{}
	val := fmt.Sprintf("%v", JsonString)
	errJson := json.Unmarshal([]byte(val), &JsonMap)
	if errJson != nil {
		return "", "", errors.New(Words.InvalidaData)
	}
	dataInJson := JsonMap[Key]
	if dataInJson == "" {
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
func IsJson(in string,attribiute string) (value interface{},valuetype string,isjson bool) {
	val := fmt.Sprintf("%v", in)
	valid := json.Valid([]byte(val))
	_,isint:=IsInt(in)
	_,isbool:=IsBool(in)
	if valid && !isint && !isbool {
		resualt,valuetype,err:=GetDataFromJsom(in,attribiute)
		if err!=nil{
			return nil,valuetype,true
		}
		return resualt,valuetype,true
	}
	return nil,valuetype,false
}
func IsLowerThan(in int, lenght int) bool {
	//if !IsInt(in) || !IsInt(lenght) {
	//	return false
	//}
	//a := fmt.Sprintf("%v", in)
	//b := fmt.Sprintf("%v", lenght)
	//
	//var x, y int
	//var err error
	//x, err = strconv.Atoi(a)
	//if err != nil {
	//	return false
	//}
	//y, err = strconv.Atoi(b)
	//if err != nil {
	//	return false
	//}
	if in < lenght {
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

func IsBetween(input int, low int, up int) bool {
/*	if !IsInt(input) || !IsInt(low) || !IsInt(up) {
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
	}*/
	if input > low && input < up {
		return true
	}
	return false
}
func IsGraterThan(in int, lenght int) bool {

/*	a := fmt.Sprintf("%v", in)
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
	}*/
	if in > lenght {
		return true
	}
	return false
}
func IsInt(input string) (out int,is bool) {
	is =false
	y, err := strconv.Atoi(input)
	if err!=nil{
		return
	}
	return y,true
}
func IsBool(in string) (value bool,is bool) {

	switch in {
	case "true":
		return true,false
	case "false":
		return false,true
	default:
		return false,false
	}
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
