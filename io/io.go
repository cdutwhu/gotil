package io

import (
	"bufio"
	"os"

	"github.com/cdutwhu/debog/base"
)

var (
	// MustWriteFile : from debog/base
	MustWriteFile = base.MustWriteFile

	// MustAppendFile : from debog/base
	MustAppendFile = base.MustAppendFile
)

// ExtractFileContent :
func ExtractFileContent(filepath string, f func(line string) (bool, string), outfile string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if ok, line := f(scanner.Text()); ok {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	content := sJoin(lines, "\n")
	if outfile != "" {
		MustWriteFile(outfile, []byte(content))
	}
	return content, nil
}
