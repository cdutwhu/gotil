package rflx

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/cdutwhu/debog/base"
	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/gotil/jx"
)

var (
	fEf    = fmt.Errorf
	fPln   = fmt.Println
	fSp    = fmt.Sprint
	fPf    = fmt.Printf
	sSplit = strings.Split

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
	// typARRAY = reflect.Array

	exist          = base.Exist
	failPOnErrWhen = fn.FailPOnErrWhen
	failPOnErr     = fn.FailPOnErr
	failOnErrWhen  = fn.FailOnErrWhen
	failOnErr      = fn.FailOnErr
	isJSON         = jx.IsJSON
)
