package rflx

import (
	"math"
	"reflect"
)

// FilterModify :
func FilterModify(slc interface{},
	filter func(i int, e interface{}) bool,
	modifier func(i int, e interface{}) interface{},
) interface{} {

	failP1OnErrWhen(slc == nil || tof(slc).Kind() != reflect.Slice, "%v: slc needs [slice]", fEf("PARAM_INVALID"))

	gSlc, r := ToGSlc(slc), []interface{}{}

	switch {
	case filter != nil && modifier != nil:
		for i, e := range gSlc {
			if filter(i, e) {
				r = append(r, modifier(i, e))
			}
		}
	case filter != nil && modifier == nil:
		for i, e := range gSlc {
			if filter(i, e) {
				r = append(r, e)
			}
		}
	case filter == nil && modifier != nil:
		for i, e := range gSlc {
			r = append(r, modifier(i, e))
		}
	default:
		r = gSlc
	}

	if len(r) > 0 {
		return ToTSlc(r)
	}
	return mkSlc(tof(slc), 0, 1).Interface()
}

// ToSet * : convert slice to set. i.e. remove duplicated items
func ToSet(slc interface{}) interface{} {
	if slc == nil {
		return nil
	}
	failP1OnErrWhen(tof(slc).Kind() != reflect.Slice, "%v: need [slice]", fEf("PARAM_INVALID"))

	v := vof(slc)
	l := v.Len()
	if l == 0 {
		return slc
	}

	set := appendX(mkSlc(tof(slc), 0, l), v.Index(0))
NEXT:
	for i := 1; i < l; i++ {
		vItem := v.Index(i)
		for j := 0; j < set.Len(); j++ {
			if deepEqual(vItem.Interface(), set.Index(j).Interface()) {
				continue NEXT
			}
			if j == set.Len()-1 { // if vItem falls down to the last set position, which means set doesn't have this item, then add it.
				set = appendX(set, vItem)
			}
		}
	}
	return set.Interface()
}

// ToGSlc * :
func ToGSlc(slc interface{}) (gSlc []interface{}) {
	if slc == nil {
		return nil
	}
	failP1OnErrWhen(tof(slc).Kind() != reflect.Slice, "%v: need [slice]", fEf("PARAM_INVALID"))

	v := vof(slc)
	l := v.Len()
	if l == 0 {
		return []interface{}{}
	}
	for i := 0; i < l; i++ {
		gSlc = append(gSlc, v.Index(i).Interface())
	}
	return
}

// ToTSlc * : support []string, []int, []float64, []bool
func ToTSlc(gSlc []interface{}) interface{} {
	if gSlc == nil || len(gSlc) == 0 {
		return nil
	}

	eleType := func(gSlc []interface{}) reflect.Kind {
		eTypes := []reflect.Kind{}
		for _, e := range gSlc {
			eTypes = append(eTypes, tof(e).Kind())
		}
		eTypes = ToSet(eTypes).([]reflect.Kind)

		switch len(eTypes) {
		case 1:
			return eTypes[0]
		case 2:
			geTypes := ToGSlc(eTypes)
			if exist(reflect.Int, geTypes...) || exist(reflect.Float64, geTypes...) {
				return reflect.Float64
			}
		}
		return reflect.Invalid
	}

	l := len(gSlc)
	switch eleType(gSlc) {
	case reflect.String:
		slc := make([]string, l)
		for i := 0; i < l; i++ {
			slc[i] = gSlc[i].(string)
		}
		return slc
	case reflect.Int:
		slc := make([]int, l)
		for i := 0; i < l; i++ {
			slc[i] = gSlc[i].(int)
		}
		return slc
	case reflect.Float64:
		slc := make([]float64, l)
		for i := 0; i < l; i++ {
			if fNum, ok := gSlc[i].(float64); ok {
				slc[i] = fNum
				continue
			} else if iNum, ok := gSlc[i].(int); ok {
				slc[i] = float64(iNum)
				continue
			} else {
				panic("Need More Number Type supported as float64")
			}
		}
		return slc
	case reflect.Bool:
		slc := make([]bool, l)
		for i := 0; i < l; i++ {
			slc[i] = gSlc[i].(bool)
		}
		return slc
	default:
		panic("Need More Slice Type in GSlc2Slc")
	}
}

// SliceAttach * : pos >= 0
func SliceAttach(s1, s2 interface{}, pos int) interface{} {
	failP1OnErrWhen(pos < 0, "%v @ pos", fEf("PARAM_INVALID"))

	v1, v2 := vof(s1), vof(s2)
	failP1OnErrWhen(v1.Kind() != reflect.Slice, "%v @ s1", fEf("PARAM_INVALID"))
	failP1OnErrWhen(v2.Kind() != reflect.Slice, "%v @ s2", fEf("PARAM_INVALID"))

	l1, l2 := v1.Len(), v2.Len()
	if l1 > 0 && l2 > 0 {
		if pos > l1 {
			return s1
		}
		lm := int(math.Max(float64(l1), float64(l2+pos)))
		v := appendSlc(v1.Slice(0, pos), v2)
		return v.Slice(0, lm).Interface()
	}
	if l1 > 0 && l2 == 0 {
		return s1
	}
	if l1 == 0 && l2 > 0 {
		return s2
	}
	return s1
}

// SliceCover * :
func SliceCover(ss ...interface{}) interface{} {
	if len(ss) == 0 {
		return nil
	}
	attached := ss[0]
	failP1OnErrWhen(vof(attached).Kind() != reflect.Slice, "%v", fEf("PARAM_INVALID"))
	for _, s := range ss[1:] {
		failP1OnErrWhen(vof(s).Kind() != reflect.Slice, "%v", fEf("PARAM_INVALID"))
		attached = SliceAttach(attached, s, 0)
	}
	return attached
}

// CanCover * : check if setA contains setB ? return the first B-Index of which item is not in setA
func CanCover(setA, setB interface{}) (bool, int) {
	if setA == nil || setB == nil {
		return false, -1
	}

	failP1OnErrWhen(tof(setA).Kind() != reflect.Slice, "%v: need [slice] @ setA", fEf("PARAM_INVALID"))
	failP1OnErrWhen(tof(setB).Kind() != reflect.Slice, "%v: need [slice] @ setB", fEf("PARAM_INVALID"))

	vA, vB := vof(setA), vof(setB)
	lA, lB := vA.Len(), vB.Len()
NEXT:
	for j := 0; j < lB; j++ {
		b := vB.Index(j).Interface()
		for i := 0; i < lA; i++ {
			if deepEqual(b, vA.Index(i).Interface()) {
				continue NEXT
			}
			// if i == lA-1 { // if b falls down to the last vA item position, which means vA doesn't have b item, return false
			// 	return false, j
			// }
		}
		return false, j
	}
	return true, -1
}

func equal(setA, setB interface{}) bool {
	okAB, _ := CanCover(setA, setB)
	okBA, _ := CanCover(setB, setA)
	return okAB && okBA
}

// Equal * : return the first comparing set index which doesn't equal previous
func Equal(setGrp ...interface{}) (bool, int, error) {
	if setGrp == nil || len(setGrp) < 2 {
		return false, -1, fEf("At least 2 sets should be input")
	}
	for i, set := range setGrp[0 : len(setGrp)-1] {
		if !equal(set, setGrp[i+1]) {
			return false, i + 1, nil
		}
	}
	return true, -1, nil
}

// intersect :
func intersect(setA, setB interface{}) interface{} {
	if setA == nil || setB == nil {
		return nil
	}

	vA, vB := vof(setA), vof(setB)
	lA, lB := vA.Len(), vB.Len()
	set := mkSlc(tof(setA), 0, lA)
NEXT:
	for j := 0; j < lB; j++ {
		b := vB.Index(j)
		for i := 0; i < lA; i++ {
			if deepEqual(b.Interface(), vA.Index(i).Interface()) {
				set = appendX(set, b)
				continue NEXT
			}
		}
	}
	return set.Interface()
}

// Intersect :
func Intersect(sets ...interface{}) interface{} {
	if len(sets) == 0 {
		return nil
	}
	intersection := sets[0]
	failP1OnErrWhen(tof(intersection).Kind() != reflect.Slice, "%v: need [slice]", fEf("PARAM_INVALID"))
	for _, s := range sets[1:] {
		failP1OnErrWhen(tof(s).Kind() != reflect.Slice, "%v: need [slice]", fEf("PARAM_INVALID"))
		intersection = intersect(intersection, s)
	}
	return intersection
}

// union :
func union(setA, setB interface{}) interface{} {
	switch {
	case setA != nil && setB == nil:
		return setA
	case setA == nil && setB != nil:
		return setB
	case setA == nil && setB == nil:
		return nil
	}

	vA, vB := vof(setA), vof(setB)
	set := mkSlc(tof(setA), 0, vA.Len()+vB.Len())
	set = appendSlc(appendSlc(set, vA), vB)
	return ToSet(set.Interface())
}

// Union :
func Union(sets ...interface{}) interface{} {
	if len(sets) == 0 {
		return nil
	}
	uni := sets[0]
	failP1OnErrWhen(tof(uni).Kind() != reflect.Slice, "%v: need [slice]", fEf("PARAM_INVALID"))
	for _, s := range sets[1:] {
		failP1OnErrWhen(tof(s).Kind() != reflect.Slice, "%v: need [slice]", fEf("PARAM_INVALID"))
		uni = union(uni, s)
	}
	return uni
}
