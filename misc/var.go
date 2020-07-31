package misc

import (
	"fmt"
	"runtime"

	"github.com/cdutwhu/debog/fn"
)

var (
	fPf  = fmt.Printf
	fPln = fmt.Println
	fEf  = fmt.Errorf

	failP1OnErrWhen = fn.FailP1OnErrWhen
	failP1OnErr     = fn.FailP1OnErr
	failOnErrWhen   = fn.FailOnErrWhen
	failOnErr       = fn.FailOnErr
)

var (
	cores = runtime.NumCPU()
)
