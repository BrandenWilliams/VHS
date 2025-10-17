package flagparse

import (
	"flag"
	"fmt"
	"os"
)

type FlagParse struct {
	In          *string
	Out         *string
	Overwrite   *bool
	PrintPreset *bool
	Preset      *int
}

func (fp *FlagParse) defineFlags() {
	fp.In = flag.String("in", "", "input file (required)")
	fp.Out = flag.String("out", "", "output file (required)")
	fp.Overwrite = flag.Bool("overwrite", false, "overwrite files(false by default)")
	fp.PrintPreset = flag.Bool("printPresets", false, "print out preset list")
	fp.Preset = flag.Int("ps", 0, "preset selection")

	flag.Usage = func() {
		printUsage()
		flag.PrintDefaults()
	}
}

func (fp *FlagParse) checkFlags() (err error) {
	if flag.NArg() > 0 {
		return fmt.Errorf("unexpected arguments: %v", flag.Args())
	}

	if *fp.PrintPreset {
		printPresets()
		os.Exit(0)
	}

	in := flag.Lookup("in").Value.String()
	out := flag.Lookup("out").Value.String()

	if in == "" || out == "" {
		err = fmt.Errorf("required flags: -in and -out")
		return
	}

	return nil
}

func (fp *FlagParse) FlagParse() (err error) {
	fp.defineFlags()

	flag.Parse()

	if err := fp.checkFlags(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		flag.Usage()
		os.Exit(2) // 2 = usage error
	}

	return
}
