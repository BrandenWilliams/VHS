package flagparse

import "fmt"

func printPresets() {
	fmt.Printf("Preset List (use number):\n" +
		"1) Libx264 (default)\n" +
		"2) VTB264\n" +
		"3) SVTAV1\n" +
		"4) Libx265\n\n")
}

func printUsage() {
	fmt.Printf("VHS Version 0.1\n" +
		"Usage:\n" +
		"	VHS -in INPUT -out OUTPUT [options]\n\n" +
		"Options:\n" +
		"	-ps PRESET\n\n" +
		"Preset List (use number):\n" +
		"1) Libx264 (default)\n" +
		"2) VTB264\n" +
		"3) SVTAV1\n" +
		"4) Libx265\n\n" +
		"Default Flags:\n")
}

func printAbsoluteDirectory() {
	fmt.Printf("use absolute directory instead of relative\n" +
		"Use DIR/ or ./DIR/")
}
