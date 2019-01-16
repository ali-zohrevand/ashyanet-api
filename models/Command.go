package models

type Command struct {
	Name  string `json:"name" valid:"runelength(1|200),blacklist~Bad Char"`
	Value string `json:"value" valid:"runelength(1|200),blacklist~Bad Char"`
	Dsc   string `json:"dsc"`
	Topic string `json:"topic" valid:"runelength(1|200),blacklist~Bad Char"`
}
type Data struct {
	Name      string `json:"name" valid:"runelength(1|200),blacklist~Bad Char"`
	ValueType string `json:"value_type" valid:"runelength(1|200),blacklist~Bad Char"`
	Dsc       string `json:"dsc"`
	Topic     string `json:"topic" valid:"runelength(1|200),blacklist~Bad Char"`
}
