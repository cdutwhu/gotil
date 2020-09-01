package embres

import (
	"testing"
)

func TestFileBytes(t *testing.T) {
	SetResAlias("Abc", "../go.sum")
	PrintFileBytes("embres", "MapRes", "", false, "../go.sum")
}

func TestCreateResDirBytesFile(t *testing.T) {
	CreateDirBytes("embres", "MapRes", "../", "./aaa.go", true)
	fPln(string(MapRes["Auto0002"]))
}
