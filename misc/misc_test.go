package misc

import (
	"testing"
	"time"
)

func TestTrackTime(t *testing.T) {
	defer TrackTime(time.Now())
	time.Sleep(time.Second * time.Duration(5))
}

func TestIF(t *testing.T) {
	a := IF(1 > 2, 1, 2).(int)
	fPln(a + 3)
}

func TestMatchAssign(t *testing.T) {
	r := MatchAssign(5, 5, 4, "a", "b", "def")
	fPln(r)
}

func f(dim, tid int, done chan int, params ...interface{}) {
	slc := params[0].([]int)
	L := len(slc)
	for i := tid; i < L; i += dim {
		time.Sleep(time.Millisecond * time.Duration(slc[i]))
	}
	done <- tid
}

func TestGo(t *testing.T) {
	delay := []int{200, 400, 600, 800, 1000, 1200, 1400, 1600}

	defer TrackTime(time.Now())
	fPln("Doing...")

	// Go(1, f, delay)
	// Go(2, f, delay)
	// Go(3, f, delay)
	// Go(4, f, delay)
	// Go(len(delay), f, delay)
	for i := 0; i < len(delay); i++ {
		time.Sleep(time.Millisecond * time.Duration(delay[i]))
	}
	fPln("OK...")
}
