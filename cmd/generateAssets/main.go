package main

import (
	"flag"
	"io/fs"
	// "os"
	"path/filepath"

	"github.com/dave/jennifer/jen"
)

func main() {
	var dirName string
	flag.StringVar(&dirName, "path", "directory", "the path to the assets directory")
	file := jen.NewFile("assetFS")

	filepath.WalkDir("dirName", func(path string, d fs.DirEntry, err error) error {
		if d.Type().IsRegular() {
			file.Func()// .Id(path)
		}
	})
}