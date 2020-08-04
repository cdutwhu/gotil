package misc

import (
	"sync"
	"time"
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

// Go : Async dispatch n threads of func. Once all done, Sync then return.
// DO NOT apply to compute with a shared variable as slow than normal.
// Start [f] implement with `defer func() { done <- tid }()`
func Go(n int, f func(dim, tid int, done chan int, params ...interface{}), params ...interface{}) {
	if n < 1 {
		n = 1
	}
	wg, done := sync.WaitGroup{}, make(chan int, n)
	wg.Add(n)
	for i := 0; i < n; i++ {
		go f(n, i, done, params...)
	}
	i := 0
	for range done {
		if i == n-1 {
			close(done)
		}
		i++
	}
	wg.Done()
}
