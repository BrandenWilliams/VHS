package ffconvert

import (
	"fmt"
	"path/filepath"
)

func getAllAbsolutePaths(inPath, outDir string) (inAbs, outAbs string, err error) {
	if inAbs, err = filepath.Abs(inPath); err != nil {
		return "", "", fmt.Errorf("absolute input error: %w", err)
	}

	if outAbs, err = filepath.Abs(outDir); err != nil {
		return "", "", fmt.Errorf("absolute outdir error: %w", err)
	}

	return
}

func generateOutPath(inAbs, destAbs string, fileExt string) string {
	base := filepath.Base(inAbs)
	ext := filepath.Ext(base)
	name := base[:len(base)-len(ext)]
	return filepath.Join(destAbs, name+fileExt)
}

func checkIfSameFilePath(a, b string) bool {
	ra, _ := filepath.EvalSymlinks(a)
	rb, _ := filepath.EvalSymlinks(b)

	if ra == "" {
		ra = a
	}

	if rb == "" {
		rb = b
	}

	return ra == rb
}
