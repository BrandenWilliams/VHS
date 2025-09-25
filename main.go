package main

import (
	"log"

	"github.com/FFNormalMovies/ffconvert"
)

func main() {
	var (
		ffc ffconvert.FFConvert

		// test vars
		inDir     string
		destDir   string
		crf       int
		preset    string
		overwrite bool
	)

	inDir = "./videos/"
	destDir = "./newvideos"
	crf = 23
	preset = ""
	overwrite = false

	ffc.NewFFConvert(inDir, destDir, crf, preset, overwrite)

	if err := ffc.ConvertVideos(); err != nil {
		log.Fatal(err)
	}
}
