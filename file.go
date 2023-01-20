package xutil

import (
	"compress/gzip"
	"io"
	"os"

	"github.com/gookit/goutil/envutil"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/strutil"
	"github.com/ungerik/go-dry"
)

// RefineString replaces chars those not allowed in Windows `\/:*?"<>|`
func RefineString(name string) string {
	g2u := map[string]string{
		"/":  "_",
		"\\": "_",
		"\"": "_",
		":":  "_",
		"*":  "_",
		"?":  "_",
		"<":  "_",
		">":  "_",
		"|":  "_",
	}
	name = strutil.Replaces(name, g2u)
	return name
}

func RefineWinFileName(name string) string {
	if !envutil.IsWindows() {
		return name
	}
	return RefineString(name)
}

func WriteFile(name string, data string) error {
	if _, err := fsutil.CreateFile(name, os.ModePerm, os.ModePerm); err != nil {
		return err
	}
	return dry.FileSetString(name, data)
}

func AddFileIfNotExisted(name, data string) string {
	name = fsutil.ExpandPath(name)
	if fsutil.FileExists(name) {
		return ""
	}
	MustWriteFile(name, data)
	return name
}

func MustWriteFile(name, data string) {
	if err := WriteFile(name, data); err != nil {
		PanicIfErr(err)
	}
}

func MkdirIfNotExist(dirPath string) (b bool, dirAbsPath string) {
	dirAbsPath = fsutil.ExpandPath(dirPath)
	if b := fsutil.PathExists(dirAbsPath); !b {
		fsutil.Mkdir(dirAbsPath, os.ModePerm)
		return true, dirAbsPath
	}
	return
}

func ReadFile(filename string) (raw []byte, err error) {
	if IsGzip(filename) {
		raw, err = FileGetGz(filename)
	} else {
		raw, err = dry.FileGetBytes(filename)
	}
	return
}

func IsGzip(filename string) bool {
	b, e := dry.FileGetBytes(filename)
	if e != nil {
		return false
	}

	if b[0] == 31 && b[1] == 139 {
		return true
	}
	return false
}

func FileGetGz(filename string) ([]byte, error) {
	fi, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fi.Close()

	fz, err := gzip.NewReader(fi)
	if err != nil {
		return nil, err
	}
	defer fz.Close()

	s, err := io.ReadAll(fz)
	if err != nil {
		return nil, err
	}
	return s, nil
}
