package misc

import "time"

var (
	// Exist : from debog/base
	Exist = exist
)

// TrackTime : defer TrackTime(time.Now())
func TrackTime(start time.Time) {
	elapsed := time.Since(start)
	fPf("Took %s\n", elapsed)
}

// IF : Ternary Operator LIKE < ? : >, BUT NO S/C, so src1 and src2 MUST all valid. e.g. type assert, nil pointer, out of index
func IF(condition bool, src1, src2 interface{}) interface{} {
	if condition {
		return src1
	}
	return src2
}

// MatchAssign : NO ShortCut, MUST all valid, e.g. type assert, nil pointer, out of index
// MatchAssign(chk, case1, case2, value1, value2, default)
func MatchAssign(chkCasesValues ...interface{}) interface{} {
	l := len(chkCasesValues)
	failP1OnErrWhen(l < 4 || l%2 == 1, "%v", fEf("PARAM_INVALID"))

	_, l1, l2 := 1, (l-1)/2, (l-1)/2
	check := chkCasesValues[0]
	cases := chkCasesValues[1 : 1+l1]
	values := chkCasesValues[1+l1 : 1+l1+l2]
	for i, c := range cases {
		if check == c {
			return values[i]
		}
	}
	return chkCasesValues[l-1]
}
