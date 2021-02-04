package io

import (
	"testing"
)

func TestIO(t *testing.T) {
	MustWriteFile("./a/b.txt", []byte("write"))
	MustAppendFile("./a/b.txt", []byte("append"), true)
}

func TestExtractFileContent(t *testing.T) {
	ExtractFileContent("/home/qmiao/Desktop/out.txt", func(ln string) (bool, string) {
		if sHasPrefix(ln, "NUMERIC:") {
			return true, ln[len("NUMERIC"):]
		}
		return false, ""
	}, "out.txt")
}
