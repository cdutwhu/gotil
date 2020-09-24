package dispatcher

import (
	"testing"
	"time"

	"github.com/cdutwhu/gotil/misc"
)

func f(dim, tid int, done chan int, params ...interface{}) {
	defer func() { done <- tid }()
	slc := params[0].([]int)
	for i := tid; i < len(slc); i += dim {
		fPf("%d : start\n", i)
		time.Sleep(time.Millisecond * time.Duration(slc[i]))
		fPf("%d : end\n", i)
	}
}

func TestSyncParallel(t *testing.T) {
	delay := []int{200, 400, 600, 800, 1000, 1200, 1400, 1600}
	defer misc.TrackTime(time.Now())
	fPln("Doing...")
	SyncParallel(4, f, delay)
	fPln("OK...")
}
