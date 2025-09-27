package ffconvert

import (
	"fmt"
	"os"
	"path/filepath"
)

func getAllAbsolutePaths(inDir, outDir string) (inAbs, outAbs string, err error) {
	if inAbs, err = filepath.Abs(inDir); err != nil {
		return "", "", fmt.Errorf("absoulte indir error: %w", err)
	}

	if outAbs, err = filepath.Abs(outDir); err != nil {
		return "", "", fmt.Errorf("absolute outdir error: %w", err)
	}

	return
}

func generateOutPathArg(currentFile, destAbs string, fileExt string) (arg string, err error) {
	if currentFile == "" {
		err = fmt.Errorf("currentfile in empty")
		return
	} else if destAbs == "" {
		err = fmt.Errorf("absolute destination empty")
		return
	} else if fileExt == "" {
		err = fmt.Errorf("filext empty")
		return
	}

	base := filepath.Base(currentFile)
	ext := filepath.Ext(base)
	name := base[:len(base)-len(ext)]
	arg = filepath.Join(destAbs, name+fileExt)
	return
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

func checkDirEntryForContinue(d os.DirEntry) bool {
	if d.IsDir() {
		return true
	}

	if d.Name() == ".DS_Store" {
		return true
	}

	return false
}
