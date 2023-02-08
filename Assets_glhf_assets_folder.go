//go:build !glhf_no_assets_folder

package glhf

import (
	"embed"
	"path"
)

//go:embed assets/*
var _glhf_assets_folder embed.FS

func init() {

	entries, err := _glhf_assets_folder.ReadDir(path.Join(".", path.Join("/", "assets")))
	if err != nil {
		panic(err)
	}
	for _, entry := range entries {
		println("loaded:", entry.Name())
	}

	err = assets.MountFS("glhf", _glhf_assets_folder)
	if err != nil {
		panic(err)
	}
}
