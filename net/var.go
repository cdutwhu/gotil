package net

import (
	"fmt"
	"strings"

	"github.com/cdutwhu/debog/fn"
)

var (
	fSf         = fmt.Sprintf
	fPln        = fmt.Println
	sHasSuffix  = strings.HasSuffix
	sIndex      = strings.Index
	sTrimRight  = strings.TrimRight
	sTrimSuffix = strings.TrimSuffix
	warnOnErr   = fn.WarnOnErr
)
