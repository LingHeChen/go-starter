package ui

import (
	"embed"
	"io/fs"
)

//go:embed build/*
var distFS embed.FS

func GetDistFS() fs.FS {
	f, err := fs.Sub(distFS, "build")
	if err != nil {
		panic(err)
	}
	return f
}
