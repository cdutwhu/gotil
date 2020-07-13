package endec

import (
	"fmt"
	"regexp"

	"github.com/cdutwhu/debog/fn"
)

var (
	fPf       = fmt.Printf
	fSf       = fmt.Sprintf
	fPln      = fmt.Println
	failOnErr = fn.FailOnErr
)

var (
	// RegexpMD5 : regular expression for MD5 string
	RegexpMD5 = regexp.MustCompile("\"[A-Fa-f0-9]{32}\"")
	// RegexpSHA1 : regular expression for SHA1 string
	RegexpSHA1 = regexp.MustCompile("\"[A-Fa-f0-9]{40}\"")
	// RegexpSHA256 : regular expression for SHA256 string
	RegexpSHA256 = regexp.MustCompile("\"[A-Fa-f0-9]{64}\"")
)
