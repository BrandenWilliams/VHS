package main

import (
	"log"

	"github.com/BrandenWilliams/VHS/ffconvert"
	"github.com/BrandenWilliams/VHS/flagparse"
)

func main() {
	var (
		fp  flagparse.FlagParse
		ffc ffconvert.FFConvert
	)

	fp.FlagParse()

	overwrite := true
	ffc.NewFFConvert(*fp.In, *fp.Out, *fp.Preset, overwrite)

	if err := ffc.FFConvert(); err != nil {
		log.Fatal(err)
	}
}
