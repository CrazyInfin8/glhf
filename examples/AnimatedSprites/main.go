package main

import (
	"embed"

	"github.com/crazyinfin8/glhf"
)

//go:embed assets/*
var _project_asset_folder embed.FS

func init() {
	err := glhf.GetAssetFS().MountFS("main", _project_asset_folder)
	if err != nil {
		panic(err)
	}
}

var g *glhf.Game

func main() {
	g = glhf.NewGame(600, 400)
	g.Start(&PlayState{})
}
