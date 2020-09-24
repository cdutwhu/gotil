package dispatcher

import (
	"time"
)

type (
	// Proc : first param is process ID
	Proc func(params ...interface{}) error

	// AsyncProc :
	AsyncProc func(deferNil chan *AsyncProc, cerr chan error, params ...interface{})

	// TaskPool :
	TaskPool struct {
		pool []AsyncProc
		fchs map[int64]chan error
	}
)

var (
	tpInit = false
)

// NewTaskPool :
func NewTaskPool(cap int) *TaskPool {
	if !tpInit {
		tpInit = true
		tp := &TaskPool{
			make([]AsyncProc, cap),
			make(map[int64]chan error),
		}
		return tp
	}
	return nil
}

func (tp *TaskPool) available(timeout int) *AsyncProc {
	try, eachWait := 0, 5
AGAIN:
	for i, t := range tp.pool {
		if t == nil {
			return &tp.pool[i]
		}
	}
	if try == timeout/eachWait {
		panic("time out - dispense")
	}
	time.Sleep(time.Millisecond * time.Duration(eachWait))
	try++
	goto AGAIN
}

// asyncProcFactory :
func (tp *TaskPool) asyncProcFactory(proc Proc, params ...interface{}) AsyncProc {
	return func(deferNil chan *AsyncProc, cerr chan error, params ...interface{}) {
		defer func() { *(<-deferNil) = nil }()
		cerr <- proc(params...)
	}
}

// AsyncTask : proc start with 'defer func() { *(<-deferNil) = nil }()'
func (tp *TaskPool) AsyncTask(pid int64, proc Proc, params ...interface{}) {
	timeout := 1000

	ptr := tp.available(timeout)
	*ptr = tp.asyncProcFactory(proc)

	deferNil := make(chan *AsyncProc)
	go func() { deferNil <- ptr }()
	time.Sleep(time.Millisecond * time.Duration(1))

	cerr := make(chan error, 1)
	tp.fchs[pid] = cerr
	// go (*ptr)(deferNil, cerr, params...)
	go (*ptr)(deferNil, cerr, append([]interface{}{pid}, params...)...)
}

// Wait :
func (tp *TaskPool) Wait() (errs map[int64]error) {
	errs = make(map[int64]error)
	for pid, fch := range tp.fchs {
		if err := <-fch; err != nil {
			errs[pid] = err
		}
	}
	if len(errs) == 0 {
		return nil
	}
	return
}
