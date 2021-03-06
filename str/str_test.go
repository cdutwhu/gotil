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
	s := "ABC0123456789ABCDEFAC"
	r := regexp.MustCompile(`AB`)
	posGrp := r.FindAllStringIndex(s, -1)
	ss := ReplByPosGrp(s, posGrp, []string{"aaa"}, 0, 1)
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

func TestSplitLn(t *testing.T) {
	txt := `		abc
777
		def`
	for _, ln := range SplitLn(txt) {
		fPln(ln)
	}
}

func TestReplAllOnAny(t *testing.T) {
	s := "a/b-c.d_e"
	fPln(s)
	s = ReplAllOnAny(s, []string{"/", "-", ".", "_"}, " ")
	fPln(s)
}

func TestExtend(t *testing.T) {
	s := "abcdefgABChijklmnABC"
	ss := sSplit(s, "ABC")
	for i, each := range ss {
		if i > 0 {
			each = Extend(each, 1, 1)
		}
		fPln(each)
	}

	s = "abcdefghijklmn"
	s = s[2:4]
	fPln(s)
	s = Extend(s, 1, 5)
	fPln(s)
	// s = Extend(s, 1, 1)
	// fPln(s)
	// s = Extend(s, 1, 2)
	// fPln(s[1:5])
}
