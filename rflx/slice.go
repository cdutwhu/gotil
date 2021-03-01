package rflx

import (
	"math"
)

// SliceAttach :
func SliceAttach(s1, s2 interface{}, pos int) interface{} {
	v1, v2 := vof(s1), vof(s2)

	k1, k2 := v1.Kind(), v2.Kind()
	failP1OnErrWhen(!(k1 == typSLICE && k2 == typSLICE), "%v", fEf("PARAM_INVALID"))

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

// SliceCover :
func SliceCover(ss ...interface{}) interface{} {
	if len(ss) == 0 {
		return nil
	}
	attached := ss[0]
	k := vof(attached).Kind()
	failP1OnErrWhen(k != typSLICE, "%v", fEf("PARAM_INVALID"))
	for _, s := range ss[1:] {
		k = vof(s).Kind()
		failP1OnErrWhen(k != typSLICE, "%v", fEf("PARAM_INVALID"))
		attached = SliceAttach(attached, s, 0)
	}
	return attached
}

// CanSetCover : check if setA contains setB ? return the first B-Index of which item is not in setA
func CanSetCover(setA, setB interface{}) (bool, int) {
	if setA == nil {
		return false, -1
	}

	tA, tB := tof(setA), tof(setB)
	kA, kB := tA.Kind(), tB.Kind()
	failP1OnErrWhen(kA != typSLICE || kB != typSLICE, "%v: need [slice]", fEf("PARAM_INVALID"))

	vA, vB := vof(setA), vof(setB)
	lA, lB := vA.Len(), vB.Len()
	if lA < lB {
		return false, -1
	}
NEXT:
	for j := 0; j < lB; j++ {
		b := vB.Index(j).Interface()
		for i := 0; i < lA; i++ {
			if deepEqual(b, vA.Index(i).Interface()) {
				continue NEXT
			}
			if i == lA-1 { // if b falls down to the last vA item position, which means vA doesn't have b item, return false
				return false, j
			}
		}
	}
	return true, -1
}

// intersect :
func intersect(setA, setB interface{}) interface{} {
	if setA == nil || setB == nil {
		return nil
	}

	tA := tof(setA)
	vA, vB := vof(setA), vof(setB)
	lA, lB := vA.Len(), vB.Len()
	set := mkSlc(tA, 0, lA)
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

// SetIntersect :
func SetIntersect(sets ...interface{}) interface{} {
	if len(sets) == 0 {
		return nil
	}
	intersection := sets[0]
	failP1OnErrWhen(tof(intersection).Kind() != typSLICE, "%v: need [slice]", fEf("PARAM_INVALID"))
	for _, s := range sets[1:] {
		failP1OnErrWhen(tof(s).Kind() != typSLICE, "%v: need [slice]", fEf("PARAM_INVALID"))
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

	tA := tof(setA)
	vA, vB := vof(setA), vof(setB)
	set := mkSlc(tA, 0, vA.Len()+vB.Len())
	set = appendSlc(appendSlc(set, vA), vB)
	return ToSet(set.Interface())
}

// SetUnion :
func SetUnion(sets ...interface{}) interface{} {
	if len(sets) == 0 {
		return nil
	}
	uni := sets[0]
	failP1OnErrWhen(tof(uni).Kind() != typSLICE, "%v: need [slice]", fEf("PARAM_INVALID"))
	for _, s := range sets[1:] {
		failP1OnErrWhen(tof(s).Kind() != typSLICE, "%v: need [slice]", fEf("PARAM_INVALID"))
		uni = union(uni, s)
	}
	return uni
}

// ToSet : convert slice to set. i.e. remove duplicated items
func ToSet(slc interface{}) interface{} {
	if slc == nil {
		return nil
	}

	t := tof(slc)
	k := t.Kind()
	failP1OnErrWhen(k != typSLICE, "%v: need [slice]", fEf("PARAM_INVALID"))

	v := vof(slc)
	l := v.Len()
	if l == 0 {
		return slc
	}

	set := mkSlc(t, 0, l)
	set = appendX(set, v.Index(0))
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

// ToGeneralSlc :
func ToGeneralSlc(slc interface{}) (gSlc []interface{}) {
	if slc == nil {
		return nil
	}

	v := vof(slc)
	k := v.Type().Kind()
	failP1OnErrWhen(k != typSLICE, "%v: need [slice]", fEf("PARAM_INVALID"))

	l := v.Len()
	for i := 0; i < l; i++ {
		gSlc = append(gSlc, v.Index(i).Interface())
	}
	return
}
