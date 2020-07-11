package misc

import (
	"fmt"

	"github.com/cdutwhu/debog/base"
	"github.com/cdutwhu/debog/fn"
)

var (
	fPf  = fmt.Printf
	fPln = fmt.Println
	fEf  = fmt.Errorf

	exist           = base.Exist
	failP1OnErrWhen = fn.FailP1OnErrWhen
	failP1OnErr     = fn.FailP1OnErr
	failOnErrWhen   = fn.FailOnErrWhen
	failOnErr       = fn.FailOnErr
)
