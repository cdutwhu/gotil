package gotil

import "github.com/cdutwhu/debog/base"

var (
	// Exist : from debog/base
	Exist = base.Exist
	// NotExist : from debog/base
	NotExist = base.NotExist
)

// IndexOf : returns the index of the first instance of e in set, or -1 if e is not present in set
func IndexOf(e interface{}, set ...interface{}) int {
	for i, ele := range set {
		if ele == e {
			return i
		}
	}
	return -1
}

// LastIndexOf : returns the index of the last instance of e in set, or -1 if e is not present in set
func LastIndexOf(e interface{}, set ...interface{}) int {
	n := len(set)
	for i := n - 1; i >= 0; i-- {
		if set[i] == e {
			return i
		}
	}
	return -1
}
