package str

import (
	"regexp"
	"testing"
)

func TestRmTailFromLast(t *testing.T) {
	fPln(RmTailFromLast("AB.CD.EF", "."))
	fPln(RmTailFromLast("AB.CD.EF", "#"))
}

func TestRmTailFromLastN(t *testing.T) {
	fPln(RmTailFromLastN("AB.CD.EF", ".", 2))
	fPln(RmTailFromLastN("AB.CD.EF", "#", 2))
}

func TestRmTailFromFirst(t *testing.T) {
	fPln(RmTailFromFirstAny(`Activity>RefId="C27E1FCF-C163-485F-BEF0-F36F18A0493A" lang="en"`, " ", ">"))
}

func TestRmHeadToLast(t *testing.T) {
	fPln(RmHeadToLast("##AB##CD##F", "D"))
}

func TestRmHeadToFirst(t *testing.T) {
	fPln(RmHeadToFirst("777##AB##CD##F", "##"))
}

func TestStrReplByPos(t *testing.T) {
	s := "ABC0123456789ABCDEF"
	r := regexp.MustCompile(`^A`)
	posGrp := r.FindAllIndex([]byte(s), -1)
	ss := ReplByPosGrp(s, posGrp, []string{"aaa"})
	fPln(ss) //         0CCC2*****BBB6789ATTTF
}

func TestTranspose(t *testing.T) {
	ss := []string{
		"12@a@1~b",
		"b~c",
		"c~a",
		"a~b~c",
		"b~c~d~e",
		"c~a",
		"d~e~b~a~f",
		"a",
		"***@b@***",
	}
	fPln(Transpose(ss, "~", "@", "@", false))
}

func TestIndentTxt(t *testing.T) {
	txt := `abc
def
ghi`
	fPln(IndentTxt(txt, 2, true))
}

func TestSplitRev(t *testing.T) {
	fPln(SplitRev("a,b,c,d,e,f,g", ","))
	fPln(SplitRev("a,b,c,d,e,f", ","))
}
