package embres

import (
	"testing"
)

func TestFileBytes(t *testing.T) {
	SetResAlias("Abc", "../README.md")
	PrintFileBytes("test", "", false, "../README.md")
}

func TestCreateResDirBytesFile(t *testing.T) {
	CreateResDirBytesFile("test", "../", "./aaa.go", true)
}
