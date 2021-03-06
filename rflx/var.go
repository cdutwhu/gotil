package rflx

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/gotil"
	"github.com/cdutwhu/gotil/judge"
)

var (
	fEf         = fmt.Errorf
	fPln        = fmt.Println
	fSp         = fmt.Sprint
	fPf         = fmt.Printf
	fPt         = fmt.Print
	fSf         = fmt.Sprintf
	sSplit      = strings.Split
	sJoin       = strings.Join
	scParseUint = strconv.ParseUint

	vof       = reflect.ValueOf
	tof       = reflect.TypeOf
	sof       = reflect.SliceOf
	mof       = reflect.MapOf
	mkSlc     = reflect.MakeSlice
	appendSlc = reflect.AppendSlice
	appendX   = reflect.Append
	mkMap     = reflect.MakeMap
	deepEqual = reflect.DeepEqual

	failPnOnErrWhen = fn.FailPnOnErrWhen
	failPnOnErr     = fn.FailPnOnErr
	failP1OnErrWhen = fn.FailP1OnErrWhen
	failP1OnErr     = fn.FailP1OnErr
	failOnErrWhen   = fn.FailOnErrWhen
	failOnErr       = fn.FailOnErr
	isJSON          = judge.IsJSON
	exist           = gotil.Exist
)

var (
	repParam = regexp.MustCompile(`^\$[0-9]+$`)
)
