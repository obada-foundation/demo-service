package web

import (
	"embed"
	"io/fs"
)

//go:embed build/*
var Assets embed.FS


// FsFunc is short-hand for constructing a http.FileSystem
// implementation
type FsFunc func(name string) (fs.File, error)

func (f FsFunc) Open(name string) (fs.File, error) {
	return f(name)
}
