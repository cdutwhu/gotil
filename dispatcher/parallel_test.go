package dispatcher

import (
	"testing"
	"time"

	"github.com/cdutwhu/gotil/misc"
)

// func f(dim, tid int, done chan int, params ...interface{}) {
// 	defer func() { done <- tid }()

// 	slc := params[0].([]int32)

//	64 BYTES
// 	L := len(slc)
// 	step := dim * 16
// 	for i := tid * 16; i < L; i += step {
// 		for j := 0; j < 16; j++ {
// 			if pos := i + j; pos < L {
// 				slc[pos]++
// 				// fPln(tid, i, j, pos, slc[pos])
// 			}
// 		}
// 	}

// 	// for i := tid; i < len(slc); i += dim {
// 	// 	slc[i]++
// 	// }
// }

// func TestSyncParallel(t *testing.T) {
// 	defer misc.TrackTime(time.Now())
// 	data := make([]int32, 100000000)
// 	fPln("Doing...")
// 	// SyncParallel(4, f, data)

// 	for i := 0; i < len(data); i++ {
// 		data[i]++
// 	}

// 	fPln("OK...")
// 	fPln(data[0], data[len(data)-1])
// }

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
	SyncParallel(5, f, delay)
	fPln("OK...")
}
