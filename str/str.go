package str

import "sort"

var (
	// RmTailFromLast : from debog/base
	RmTailFromLast = rmTailFromLast

	// RmHeadToLast : from debog/base
	RmHeadToLast = rmHeadToLast
)

// RmTailFromLastN :
func RmTailFromLastN(s, mark string, iLast int) string {
	rt := s
	for i := 0; i < iLast; i++ {
		rt = RmTailFromLast(rt, mark)
	}
	return rt
}

// RmTailFromFirst :
func RmTailFromFirst(s, mark string) string {
	if i := sIndex(s, mark); i >= 0 {
		return s[:i]
	}
	return s
}

// RmTailFromFirstAny :
func RmTailFromFirstAny(s string, marks ...string) string {
	if len(marks) == 0 {
		return s
	}
	const MAX = 1000000
	var I int = MAX
	for _, mark := range marks {
		if i := sIndex(s, mark); i >= 0 && i < I {
			I = i
		}
	}
	if I != MAX {
		return s[:I]
	}
	return s
}

// RmHeadToFirst :
func RmHeadToFirst(s, mark string) string {
	if segs := sSplit(s, mark); len(segs) > 1 {
		return sJoin(segs[1:], mark)
	}
	return s
}

// ------------------------------------------------ //

// HasAnyPrefix : [lsPrefix] at least input one prefix
func HasAnyPrefix(s string, lsPrefix ...string) bool {
	for _, prefix := range lsPrefix {
		if sHasPrefix(s, prefix) {
			return true
		}
	}
	return false
}

// HasAnySuffix : [lsSuffix] at least input one suffix
func HasAnySuffix(s string, lsSuffix ...string) bool {
	for _, suffix := range lsSuffix {
		if sHasSuffix(s, suffix) {
			return true
		}
	}
	return false
}

// ReplByPosGrp : [posGrp]-[newStrGrp] same length OR newStrGrp can only have 1 element for filling into all posGrp
func ReplByPosGrp(s string, posGrp [][]int, newStrGrp []string) string {
	if len(posGrp) == 0 {
		return s
	}
	failP1OnErrWhen(!(len(posGrp) == len(newStrGrp) || len(newStrGrp) == 1), "%v", fEf("SLICE_INCORRECT_COUNT"))

	wrapper := make([]struct {
		posPair []int
		newStr  string
	}, len(posGrp))
	for i, pair := range posGrp {
		wrapper[i].posPair = pair
		if len(newStrGrp) == 1 {
			wrapper[i].newStr = newStrGrp[0]
		} else {
			wrapper[i].newStr = newStrGrp[i]
		}
	}
	sort.Slice(wrapper, func(i, j int) bool {
		return wrapper[i].posPair[0] < wrapper[j].posPair[0]
	})

	complement := make([][2]int, len(posGrp)+1)
	for i := 0; i < len(complement); i++ {
		if i == 0 {
			complement[i][0] = 0
			complement[i][1] = wrapper[i].posPair[0]
		} else if i == len(complement)-1 {
			complement[i][0] = wrapper[i-1].posPair[1]
			complement[i][1] = len(s)
		} else {
			complement[i][0] = wrapper[i-1].posPair[1]
			complement[i][1] = wrapper[i].posPair[0]
		}
	}

	keepStrGrp := make([]string, len(complement))
	for i := 0; i < len(keepStrGrp); i++ {
		start, end := complement[i][0], complement[i][1]
		failOnErrWhen(end < start, "%v: [end] must greater than [start]", fEf("VAR_INVALID"))
		keepStrGrp[i] = s[start:end]
	}

	ret := ""
	for i, keep := range keepStrGrp[:len(keepStrGrp)-1] {
		ret += (keep + wrapper[i].newStr)
	}
	ret += keepStrGrp[len(keepStrGrp)-1]
	return ret
}

// Transpose :
func Transpose(strlist []string, sep, trimToL, trimFromR string, toset bool) [][]string {
	nSep := 0
	for _, str := range strlist {
		if n := sCount(str, sep); n > nSep {
			nSep = n
		}
	}
	rtStrList := make([][]string, nSep+1)
	for _, str := range strlist {
		for i, s := range sSplit(str, sep) {
			if trimToL != "" {
				if fd := sIndex(s, trimToL); fd >= 0 {
					s = s[fd+1:]
				}
			}
			if trimFromR != "" {
				if fd := sLastIndex(s, trimFromR); fd >= 0 {
					s = s[:fd]
				}
			}
			rtStrList[i] = append(rtStrList[i], s)
		}
	}
	if toset {
		for i := 0; i < len(rtStrList); i++ {
			rtStrList[i] = toSet(rtStrList[i]).([]string)
		}
	}
	return rtStrList
}

// IndentTxt :
func IndentTxt(str string, n int, ignore1stln bool) string {
	if n == 0 {
		return str
	}
	S := 0
	if ignore1stln {
		S = 1
	}
	lines := sSplit(str, "\n")
	if n > 0 {
		space := ""
		for i := 0; i < n; i++ {
			space += " "
		}
		for i := S; i < len(lines); i++ {
			if sTrim(lines[i], " \n\t") != "" {
				lines[i] = space + lines[i]
			}
		}
	} else {
		for i := S; i < len(lines); i++ {
			if len(lines[i]) == 0 { //                                         ignore empty string line
				continue
			}
			if len(lines[i]) <= -n || sTrimLeft(lines[i][:-n], " ") != "" { // cannot be indented as <n>, give up indent
				return str
			}
			lines[i] = lines[i][-n:]
		}
	}
	return sJoin(lines, "\n")
}
