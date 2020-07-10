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

	failOnErr  = fn.FailOnErr
	failPOnErr = fn.FailPOnErr
)
