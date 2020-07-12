package rflx

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/cdutwhu/debog/base"
	"github.com/cdutwhu/debog/fn"
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
	mkslc     = reflect.MakeSlice
	appendslc = reflect.AppendSlice
	appendx   = reflect.Append
	mkmap     = reflect.MakeMap
	deepEqual = reflect.DeepEqual

	typMAP    = reflect.Map
	typSLICE  = reflect.Slice
	typSTRUCT = reflect.Struct
	typPTR    = reflect.Ptr
	typARRAY  = reflect.Array

	exist           = base.Exist
	failPnOnErrWhen = fn.FailPnOnErrWhen
	failPnOnErr     = fn.FailPnOnErr
	failP1OnErrWhen = fn.FailP1OnErrWhen
	failP1OnErr     = fn.FailP1OnErr
	failOnErrWhen   = fn.FailOnErrWhen
	failOnErr       = fn.FailOnErr
	isJSON          = judge.IsJSON
)

var (
	repParam = regexp.MustCompile(`^\$[0-9]+$`)
)
