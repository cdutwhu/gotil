package io

import (
	"testing"
)

func TestIO(t *testing.T) {
	MustWriteFile("./a/b.txt", []byte("write"))
	MustAppendFile("./a/b.txt", []byte("append"), true)
}
