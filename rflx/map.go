package rflx

import (
	"reflect"
	"regexp"
	"sort"
)

// MapKeys : only apply to single type key
func MapKeys(m interface{}) interface{} {
	v := vof(m)
	failP1OnErrWhen(v.Kind() != reflect.Map, "%v", fEf("PARAM_INVALID_MAP"))

	keys := v.MapKeys()
	if L := len(keys); L > 0 {
		kType := tof(keys[0].Interface())
		rstValue := mkSlc(sof(kType), L, L)
		for i, k := range keys {
			rstValue.Index(i).Set(vof(k.Interface()))
		}
		// sort keys if keys are int or float64 or string
		rst := rstValue.Interface()
		switch keys[0].Interface().(type) {
		case int:
			sort.Ints(rst.([]int))
		case float64:
			sort.Float64s(rst.([]float64))
		case string:
			sort.Strings(rst.([]string))
		}
		return rst
	}
	return nil
}

// MapKVs : only apply to single type key and single type value
func MapKVs(m interface{}) (interface{}, interface{}) {
	v := vof(m)
	failP1OnErrWhen(v.Kind() != reflect.Map, "%v", fEf("PARAM_INVALID_MAP"))

	keys := v.MapKeys()
	if L := len(keys); L > 0 {
		kType := tof(keys[0].Interface())
		kRst := mkSlc(sof(kType), L, L)
		vType := tof(v.MapIndex(keys[0]).Interface())
		vRst := mkSlc(sof(vType), L, L)
		for i, k := range keys {
			kRst.Index(i).Set(vof(k.Interface()))
			vRst.Index(i).Set(vof(v.MapIndex(k).Interface()))
		}
		return kRst.Interface(), vRst.Interface()
	}
	return nil, nil
}

// mapJoin : overwritted by the 2nd params
func mapJoin(m1, m2 interface{}) (interface{}, error) {
	v1, v2 := vof(m1), vof(m2)
	// failP1OnErrWhen(v1.Kind() != reflect.Map || v2.Kind() != reflect.Map, "%v", fEf("PARAM_INVALID_MAP"))
	if v1.Kind() != reflect.Map || v2.Kind() != reflect.Map {
		return nil, fEf("PARAM_INVALID_MAP")
	}

	keys1, keys2 := v1.MapKeys(), v2.MapKeys()
	if len(keys1) > 0 && len(keys2) > 0 {
		k1, k2 := keys1[0], keys2[0]
		k1Type, k2Type := tof(k1.Interface()), tof(k2.Interface())
		v1Type, v2Type := tof(v1.MapIndex(k1).Interface()), tof(v2.MapIndex(k2).Interface())
		// failP1OnErrWhen(k1Type != k2Type, "%v", fEf("MAPS_DIF_KEY_TYPE"))
		if k1Type != k2Type {
			return nil, fEf("MAPS_DIF_KEY_TYPE")
		}
		// failP1OnErrWhen(v1Type != v2Type, "%v", fEf("MAPS_DIF_VALUE_TYPE"))
		if v1Type != v2Type {
			return nil, fEf("MAPS_DIF_VALUE_TYPE")
		}

		aMap := mkMap(mof(k1Type, v1Type))
		for _, k := range keys1 {
			aMap.SetMapIndex(vof(k.Interface()), vof(v1.MapIndex(k).Interface()))
		}
		for _, k := range keys2 {
			aMap.SetMapIndex(vof(k.Interface()), vof(v2.MapIndex(k).Interface()))
		}
		return aMap.Interface(), nil
	}
	if len(keys1) > 0 && len(keys2) == 0 {
		return m1, nil
	}
	if len(keys1) == 0 && len(keys2) > 0 {
		return m2, nil
	}
	return m1, nil
}

// MapMerge : overwritted by the later maps
func MapMerge(ms ...interface{}) interface{} {
	if len(ms) == 0 {
		return nil
	}
	var e error
	mm := ms[0]
	for _, m := range ms[1:] {
		mm, e = mapJoin(mm, m)
		failP1OnErr("%v", e)
	}
	return mm
}

// KeySortPrint : Key Sorted Print
func KeySortPrint(m interface{}) {
	re := regexp.MustCompile(`^[+-]?[0-9]*\.?[0-9]+:`)
	mapstr := fSp(m)
	mapstr = mapstr[4 : len(mapstr)-1]
	// fPln(mapstr)
	I := 0
	rmIdxList := []interface{}{}
	ss := sSplit(mapstr, " ")
	for i, s := range ss {
		if re.MatchString(s) {
			I = i
		} else {
			ss[I] += " " + s
			rmIdxList = append(rmIdxList, i) // to be deleted (i)
		}
	}
	for i, s := range ss {
		if !exist(i, rmIdxList...) {
			fPln(s)
		}
	}
}
