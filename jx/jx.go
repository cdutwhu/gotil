package jx

import (
	"encoding/json"
	"encoding/xml"
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
