package dispatcher

import (
	"fmt"
	"testing"
	"time"

	"github.com/cdutwhu/gotil/misc"
)

func TestTaskPool(t *testing.T) {
	// numcpu := flag.Int("cpu", runtime.NumCPU(), "")
	// flag.Parse()
	// fmt.Println(*numcpu)
	// runtime.GOMAXPROCS(*numcpu)

	defer misc.TrackTime(time.Now())

	// params[0] is process ID
	f := func(params ...interface{}) error {
		fmt.Printf("%d : start\n", params[0])
		time.Sleep(time.Millisecond * time.Duration(200))
		fmt.Printf("%d : end\n", params[0])
		return nil
	}

	f1 := func(params ...interface{}) error {
		fmt.Printf("%d : start\n", params[0])
		time.Sleep(time.Millisecond * time.Duration(2000))
		fmt.Printf("%d : end\n", params[0])
		return fmt.Errorf("error test")
	}

	tPool := NewTaskPool(8)
	tPool.AsyncTask(1, f1)
	tPool.AsyncTask(2, f)
	tPool.AsyncTask(3, f)
	tPool.AsyncTask(4, f)
	tPool.AsyncTask(5, f)
	tPool.AsyncTask(6, f)
	tPool.AsyncTask(7, f)
	tPool.AsyncTask(8, f1)
	for i, e := range tPool.Wait() {
		fmt.Println(i, e)
	}
}
