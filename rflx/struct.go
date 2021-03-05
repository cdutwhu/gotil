package rflx

import (
	"encoding/json"
	"os"
	"reflect"
)

// Struct2Env :
func Struct2Env(key string, s interface{}) {
	stVal := vof(s)
	failP1OnErrWhen(stVal.Kind() != reflect.Ptr || stVal.Elem().Kind() != reflect.Struct, "%v", fEf("PARAM_INVALID_STRUCT_PTR"))

	bytes, err := json.Marshal(s)
	failOnErr("%v", err)
	failOnErr("%v", os.Setenv(key, string(bytes)))
}

// Env2Struct :
func Env2Struct(key string, s interface{}) interface{} {
	stVal := vof(s)
	failP1OnErrWhen(stVal.Kind() != reflect.Ptr || stVal.Elem().Kind() != reflect.Struct, "%v", fEf("PARAM_INVALID_STRUCT_PTR"))

	jsonstr := os.Getenv(key)
	failP1OnErrWhen(!isJSON(jsonstr), "%v", fEf("JSON_INVALID"))
	failOnErr("%v", json.Unmarshal([]byte(jsonstr), s))
	return s
}

// Struct2Map : each field name MUST be Exportable
func Struct2Map(s interface{}) map[string]interface{} {
	stVal := vof(s)
	failP1OnErrWhen(stVal.Kind() != reflect.Ptr || stVal.Elem().Kind() != reflect.Struct, "%v", fEf("PARAM_INVALID_STRUCT_PTR"))

	ret := make(map[string]interface{})
	stValElem := stVal.Elem()
	valTyp := stValElem.Type()
	for i := 0; i < stValElem.NumField(); i++ {
		if name, field := valTyp.Field(i).Name, stValElem.Field(i); field.CanInterface() {
			ret[name] = field.Interface()

			// --------------- //
			if field.Type().Kind() == reflect.Func {
				fPln("func variable: " + name)
			}
			// --------------- //
		}
	}
	return ret
}

// StructFields :
func StructFields(s interface{}) []string {
	stVal := vof(s)
	failP1OnErrWhen(stVal.Kind() != reflect.Ptr || stVal.Elem().Kind() != reflect.Struct, "%v", fEf("PARAM_INVALID_STRUCT_PTR"))
	return MapKeys(Struct2Map(s)).([]string)
}
