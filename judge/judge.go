package judge

import (
	"encoding/json"
	"encoding/xml"
	"strconv"
)

// IsXML :
func IsXML(str string) bool {
	return xml.Unmarshal([]byte(str), new(interface{})) == nil
}

// IsJSON :
func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

// IsNumeric :
func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}