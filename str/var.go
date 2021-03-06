package str

import (
	"fmt"
	"strings"

	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/gotil/misc"
	"github.com/cdutwhu/gotil/rflx"
)

var (
	fPln            = fmt.Println
	fEf             = fmt.Errorf
	sJoin           = strings.Join
	sSplit          = strings.Split
	sIndex          = strings.Index
	sHasPrefix      = strings.HasPrefix
	sHasSuffix      = strings.HasSuffix
	sCount          = strings.Count
	sLastIndex      = strings.LastIndex
	sTrim           = strings.Trim
	sTrimLeft       = strings.TrimLeft
	sReplaceAll     = strings.ReplaceAll
	failP1OnErrWhen = fn.FailP1OnErrWhen
	failP1OnErr     = fn.FailP1OnErr
	failOnErrWhen   = fn.FailOnErrWhen
	failOnErr       = fn.FailOnErr
	cvtToSet        = rflx.ToSet
	matchAssign     = misc.MatchAssign
)
