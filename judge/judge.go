package judge

import (
	"encoding/json"
	"encoding/xml"
	"math"
	"reflect"
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

// IsContInts : check ints is continuous int slice
func IsContInts(ints []int) (ok bool, min int, max int) {
	if ints == nil || len(ints) == 0 {
		return false, math.MinInt32, math.MaxInt32
	}
	if len(ints) == 1 {
		return true, ints[0], ints[0]
	}

	s, e := ints[0], ints[len(ints)-1]
	if s < e {
		return reflect.DeepEqual(iter2slc(s, e+1), ints), s, e
	}
	return reflect.DeepEqual(iter2slc(s, e-1), ints), e, s
}
