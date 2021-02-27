package io

import (
	"testing"
)

func TestIO(t *testing.T) {
	MustWriteFile("./a/b.txt", []byte("write"))
	MustAppendFile("./a/b.txt", []byte("append"), true)
}

func TestEditFileByLine(t *testing.T) {
	EditFileByLine("/home/qmiao/Desktop/out.txt", func(ln string) (bool, string) {
		// if sHasPrefix(ln, "NUMERIC:") {
		// 	return true, ln[len("NUMERIC: "):]
		// }
		// if sHasPrefix(ln, "BOOLEAN:") {
		// 	return true, ln[len("BOOLEAN: "):]
		// }
		if sHasPrefix(ln, "LIST:") {
			part := ln[len("LIST: "):]
			ss := sSplit(part, "/")
			return true, sJoin(ss[:len(ss)-1], "/")
		}
		return false, ""
	}, "LIST.txt")
}

func TestEditStrByLine(t *testing.T) {
	teststr := `hello
	world
	Hello
	World`
	out, _ := EditStrByLine(teststr, func(ln string) (bool, string) {
		ln = sTrim(ln, " \t")
		if sHasPrefix(ln, "H") || sHasPrefix(ln, "W") {
			return true, ln
		}
		return false, ""
	}, "hello.txt")
	fPln(out)
}
