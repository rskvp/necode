package openapi

import (
	"embed"
	"io/fs"
)

//go:embed
var assets embed.FS

func Assets() (fs.FS, error) {
	return fs.Sub(assets, "assets")
}
