package iter

import (
	"fmt"

	"github.com/cdutwhu/debog/fn"
)

var (
	fPf  = fmt.Printf
	fPln = fmt.Println
	fSf  = fmt.Sprintf
	fEf  = fmt.Errorf

	failOnErr       = fn.FailOnErr
	failP1OnErr     = fn.FailP1OnErr
	failOnErrWhen   = fn.FailOnErrWhen
	failP1OnErrWhen = fn.FailP1OnErrWhen
)
