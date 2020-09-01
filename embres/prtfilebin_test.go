package embres

import (
	"testing"
)

func TestFileBytes(t *testing.T) {
	SetResAlias("Abc", "../go.sum")
	PrintFileBytes("embres", "MapRes", "./test.go", false, "../go.sum")
	// fPln(string(MapRes["Abc"]))
}

func TestCreateResDirBytesFile(t *testing.T) {
	CreateDirBytes("embres", "MapRes", "../", "./aaa.go", true, "Git")
	// fPln(string(MapRes["auto0002"]))
}
