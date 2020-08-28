package embres

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"
)

var (
	mPathAlias   = make(map[string][]string)
	mAliasPath   = make(map[string]string)
	rNameRule    = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)
	rPkgNameRule = regexp.MustCompile(`^[a-z][a-z0-9]*$`)
)

// SetResAlias :
func SetResAlias(alias, file string) {
	_, err := os.Stat(file)
	failP1OnErr("%v", err)
	fpAbs, err := filepath.Abs(file)
	failP1OnErr("%v", err)

	failP1OnErrWhen(!rNameRule.MatchString(alias), "%v", fEf("[%s] error: Must Follow Variable Naming Rule", alias))
	warnP1OnErrWhen(unicode.IsLower(rune(alias[0])) || alias[0] == '_', "%v", fEf("[%s] is NOT Exportable for [%s]", alias, file))

	_, exist := mAliasPath[alias]
	failP1OnErrWhen(exist, "%v", fEf("alias [%s] is already used by [%s]", alias, mAliasPath[alias]))
	mAliasPath[alias], mPathAlias[fpAbs] = fpAbs, append(mPathAlias[fpAbs], alias)
}

// PrintFileBytes :
func PrintFileBytes(pkg, savepath string, save bool, files ...string) string {
	failP1OnErrWhen(!rPkgNameRule.MatchString(pkg), "%v", fEf("[%s] error: Must Follow Package Naming Rule", pkg))

	sb := strings.Builder{}
	for _, file := range files {
		fpAbs, err := filepath.Abs(file)
		failP1OnErr("%v", err)

		bytes, err := ioutil.ReadFile(file)
		failP1OnErr("%v", err)

		if _, exist := mPathAlias[fpAbs]; !exist {
			bytesName := rmTailFromLast(file, ".")
			bytesName = replAllOnAny(bytesName, []string{".", "-"}, "")
			bytesName = replAllOnAny(bytesName, []string{"/"}, "_")
			bytesName = sTrimLeft(bytesName, " \t_") + "_" + sTrimLeft(filepath.Ext(file), ".")
			mPathAlias[fpAbs] = append([]string{}, sTitle(bytesName))
		}

		for _, alias := range mPathAlias[fpAbs] {
			failP1OnErrWhen(!rNameRule.MatchString(alias), "%v", fEf("[%s] error: Must Follow Variable Naming Rule", alias))

			sb.WriteString(fSf("var %s = [...]byte{\n\t", alias))
			for i, v := range bytes {
				if i > 0 {
					if i%32 == 0 {
						sb.WriteString(",\n\t")
					} else {
						sb.WriteString(", ")
					}
				}
				sb.WriteString(fSf("0X%02x", v))
			}
			sb.WriteString(",\n}")
			sb.WriteString("\n\n")
		}
	}

	outdir := fSf("./cache/%s/", pkg)
	if save {
		outdir = savepath
	} else {
		os.RemoveAll(outdir) // delete old package
		savepath = outdir + pkg + ".go"
	}
	content := fSf("package %s\n\n", pkg) + sb.String()
	mustWriteFile(savepath, []byte(content))
	return content
}

// CreateResDirBytesFile :
func CreateResDirBytesFile(pkg, dir, savepath string, save bool) {
	fdir, err := os.Open(dir)
	failP1OnErr("%v", err)
	dInfo, err := fdir.Stat()
	failP1OnErr("%v", err)
	failP1OnErrWhen(!dInfo.IsDir(), "%v", fEf("input dir is invalid directory"))

	resGrp := []string{}
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && !strings.HasSuffix(info.Name(), ".go") {
			resGrp = append(resGrp, path)
		}
		return nil
	})
	failP1OnErr("%v", err)
	PrintFileBytes(pkg, savepath, save, resGrp...)
}
