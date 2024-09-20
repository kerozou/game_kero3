package kero3

import (
	"embed"
	"io/fs"
	"log"
)

//go:embed assets
var rawEmbed embed.FS

var (
	Embed fs.FS
	// Images fs.FS
	// Fonts fs.FS
	// Bgms fs.FS
	// Ses fs.FS
)

func init() {
	var err error
	Embed, err = fs.Sub(rawEmbed, "assets")
	if err != nil {
		log.Fatal(err)
	}
}
