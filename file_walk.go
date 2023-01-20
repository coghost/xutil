package xutil

import (
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/thoas/go-funk"
)

type WArgs struct {
	Root      string
	FilterDir string

	NameInclude string
	NameExclude string

	Exts []string
	Dirs []string
}

type WArgFunc func(o *WArgs)

func WKv(key string, val interface{}) WArgFunc {
	return func(o *WArgs) {
		funk.Set(o, val, key)
	}
}

func WithInclude(s string) WArgFunc {
	return func(o *WArgs) {
		o.NameInclude = s
	}
}

func WExts(val []string) WArgFunc {
	return func(o *WArgs) {
		o.Exts = val
	}
}

func WLimitDir(val string) WArgFunc {
	return func(o *WArgs) {
		o.FilterDir = val
	}
}

func BindOpts(opt *WArgs, opts ...WArgFunc) {
	for _, f := range opts {
		f(opt)
	}
}

func Walk(root string, opts ...WArgFunc) []string {
	opt := &WArgs{FilterDir: ""}
	BindOpts(opt, opts...)
	root = filepath.Join(root, opt.FilterDir)
	configs := walkForAllFiles(root, opts...)
	return configs
}

func walkForAllFiles(root string, opts ...WArgFunc) (dat []string) {
	filepath.Walk(
		root,
		func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if foundFile(path, opts...) {
				dat = append(dat, path)
			}
			return nil
		},
	)
	return
}

func foundFile(path string, opts ...WArgFunc) bool {
	opt := &WArgs{
		NameInclude: "",
		NameExclude: "",
		Exts:        []string{},
		Dirs:        []string{},
	}
	BindOpts(opt, opts...)
	dir, name := filepath.Split(path)

	findDir := false
	for _, d := range opt.Dirs {
		if strings.Contains(dir, d) {
			findDir = true
			break
		}
	}
	if funk.IsEmpty(opt.Dirs) {
		findDir = true
	}

	findFile := funk.Contains(opt.Exts, filepath.Ext(name))

	if opt.NameInclude != "" {
		findFile = findFile && strings.Contains(path, opt.NameInclude)
	}

	if opt.NameExclude != "" {
		findFile = findFile && !strings.Contains(path, opt.NameExclude)
	}

	return findDir && findFile
}
