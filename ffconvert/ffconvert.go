package ffconvert

import (
	"fmt"
	"os"
	"os/exec"
)

type FFConvert struct {
	InDir  string
	OutDir string

	InAbsDir  string
	OutAbsDir string

	CurrentFile string

	Crf       int
	Preset    string
	Overwrite bool

	CLIArgs []string
}

func (ffc *FFConvert) NewFFConvert(inDir, outDir string, crf int, preset string, overwrite bool) {
	ffc.InDir = inDir
	ffc.OutDir = outDir
	ffc.Crf = crf
	ffc.Preset = preset
	ffc.Overwrite = overwrite
}

// NEED TO FULLY EXPAND THIS
func (ffc *FFConvert) BuildCLIArgs() {
	ffc.CLIArgs = []string{
		"-i", ffc.InAbsDir,
		"-c:v", "libx264",
		"-preset", ffc.Preset,
		"-crf", fmt.Sprint(ffc.Crf),
		"-c:a", "aac",
		"-b:a", "160k",
	}

	if ffc.Overwrite {
		ffc.CLIArgs = append([]string{"-y"}, ffc.CLIArgs...)
	} else {
		ffc.CLIArgs = append([]string{"-n"}, ffc.CLIArgs...)
	}

	ffc.CLIArgs = append(ffc.CLIArgs, ffc.OutAbsDir)
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

func (ffc *FFConvert) ConvertVideos() (err error) {
	var dir []os.DirEntry

	if dir, err = os.ReadDir(ffc.InDir); err != nil {
		return err
	}

	for _, d := range dir {
		if cont := checkDirEntryForContinue(d); cont {
			continue
		}

		if err = ffc.ConvertVideo(d.Name()); err != nil {
			return err
		}
	}

	return nil
}

func (ffc *FFConvert) ConvertVideo(currentFile string) (err error) {
	if ffc.InAbsDir, ffc.OutAbsDir, err = getAllAbsolutePaths(ffc.InDir+currentFile, ffc.OutDir); err != nil {
		return fmt.Errorf("error within getAllAbsolutePaths(): %s", ffc.InDir)
	}

	ffc.OutAbsDir = generateOutPath(currentFile, ffc.OutAbsDir, ".mp4")

	if checkIfSameFilePath(ffc.InAbsDir, ffc.OutAbsDir) {
		return fmt.Errorf("input and output resolve to the same file: %s", ffc.InAbsDir)
	}

	ffc.BuildCLIArgs()

	cmd := exec.Command("ffmpeg", ffc.CLIArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ffmpeg failed: %w", err)
	}

	return nil
}
