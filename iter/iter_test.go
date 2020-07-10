package iter

import "testing"

func TestIter(t *testing.T) {
	for i := range N(10) {
		fPln(i)
	}
	fPln(" ------------ ")
	for i := range Iter(2, 2, 10) {
		fPln(i)
	}
	fPln(" ------------ ")
	func(slc ...int) {
		for _, a := range slc {
			fPln(a)
		}
	}(Iter2Slc(2, 4, 10)...)
}
