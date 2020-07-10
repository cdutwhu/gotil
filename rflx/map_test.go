package rflx

import "testing"

var (
	m = map[string]string{
		"0": "00",
		"3": "33",
		"a": "aa",
		"7": "77",
		"1": "11",
		"z": "zz",
	}

	m1 = map[string]string{
		"b": "bb",
	}

	m2 = map[int]string{
		1: "bb",
	}

	m3 = map[string]int{
		"1": 111,
	}
)

func TestMapKeys(t *testing.T) {
	keys := MapKeys(m)
	fPln(keys)
}

func TestMapKVs(t *testing.T) {
	ks, vs := MapKVs(m)
	fPln(ks)
	fPln(vs)
}

func TestMapsMerge(t *testing.T) {
	mm := MapMerge(m, m1, m2)
	fPln(mm)
}

func TestMapPrint(t *testing.T) {
	KeySortPrint(m)
}
