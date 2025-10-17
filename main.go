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

	if err := fp.FlagParse(); err != nil {
		log.Fatal(err)
	}

	if err := ffc.FFConvert(*fp.In, *fp.Out, *fp.Preset, *fp.Overwrite); err != nil {
		log.Fatal(err)
	}
}
