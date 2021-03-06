package embres

import (
	"fmt"
	"strings"

	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/gotil/io"
	"github.com/cdutwhu/gotil/str"
)

var (
	fPln            = fmt.Println
	fSpt            = fmt.Sprint
	fSf             = fmt.Sprintf
	fEf             = fmt.Errorf
	sReplace        = strings.Replace
	sReplaceAll     = strings.ReplaceAll
	sTitle          = strings.Title
	sTrim           = strings.Trim
	sTrimLeft       = strings.TrimLeft
	sTrimSuffix     = strings.TrimSuffix
	sContains       = strings.Contains
	failOnErr       = fn.FailOnErr
	failP1OnErr     = fn.FailP1OnErr
	failP1OnErrWhen = fn.FailP1OnErrWhen
	warnP1OnErrWhen = fn.WarnP1OnErrWhen
	mustAppendFile  = io.MustAppendFile
	mustWriteFile   = io.MustWriteFile
	rmTailFromLast  = str.RmTailFromLast
	replAllOnAny    = str.ReplAllOnAny
)
