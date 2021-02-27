package io

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/cdutwhu/debog/base"
)

var (
	// MustWriteFile : from debog/base
	MustWriteFile = base.MustWriteFile

	// MustAppendFile : from debog/base
	MustAppendFile = base.MustAppendFile
)

// readByLine :
func readByLine(r io.Reader, f func(line string) (bool, string), outfile string) (string, error) {
	lines := []string{}
	scanner := bufio.NewScanner(r)
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

// EditFileByLine :
func EditFileByLine(filepath string, f func(line string) (bool, string), outfile string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	return readByLine(file, f, outfile)
}

// EditStrByLine :
func EditStrByLine(str string, f func(line string) (bool, string), outfile string) (string, error) {
	return readByLine(strings.NewReader(str), f, outfile)
}
