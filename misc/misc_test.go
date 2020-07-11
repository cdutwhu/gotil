package misc

import (
	"testing"
	"time"
)

func TestExist(t *testing.T) {
	fPln(Exist(1, 1, 2, 3))
}

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
