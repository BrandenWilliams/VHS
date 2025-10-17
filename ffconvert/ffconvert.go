package ffconvert

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/BrandenWilliams/VHS/ffconvert/linuxcliargs"
)

type FFConvert struct {
	LCliA     linuxcliargs.LinuxCLICfg
	BuildArgs []string

	InDir  string
	OutDir string

	InAbsDir  string
	OutAbsDir string

	CurrentFile string

	Crf       int
	Overwrite bool
}

func (ffc *FFConvert) newFFConvert(inDir, outDir string, preMadeArg int, overwrite bool) {
	ffc.InDir = inDir
	ffc.OutDir = outDir
	ffc.Overwrite = overwrite

	ffc.LCliA = ffc.LCliA.SetPreMadeArg(preMadeArg)
}

func runFFMpegCommand(buildArgs []string) (err error) {
	cmd := exec.Command("ffmpeg", buildArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err = cmd.Run(); err != nil {
		return fmt.Errorf("cmdrun() ffmpeg failed: %w", err)
	}

	return nil
}

func (ffc *FFConvert) ConvertFolderVideoPrep(currentFile string) (err error) {
	if ffc.InAbsDir, ffc.OutAbsDir, err = getAllAbsolutePaths(ffc.InDir+currentFile, ffc.OutDir); err != nil {
		return fmt.Errorf("error within getallabsolutepaths(): %s", ffc.InDir)
	}

	if ffc.OutAbsDir, err = generateOutPathArg(currentFile, ffc.OutAbsDir, ".mp4"); err != nil {
		return fmt.Errorf("error within generateoutpatharg(): %s", err)
	}

	return nil
}

func (ffc *FFConvert) ConvertVideoPrep() (err error) {
	ffc.InAbsDir = ffc.InDir
	ffc.OutAbsDir = ffc.OutDir

	return
}

func (ffc *FFConvert) ConvertVideo() (err error) {
	if checkIfSameFilePath(ffc.InAbsDir, ffc.OutAbsDir) {
		return fmt.Errorf("input file and output file resolve to the same file: %s", ffc.InAbsDir)
	}

	if ffc.BuildArgs, err = ffc.LCliA.BuildLinuxCLIArgs(ffc.InAbsDir, ffc.OutAbsDir, ffc.Overwrite); err != nil {
		return fmt.Errorf("error within buildlinuxcliargs(): %s", err)
	}

	if err = runFFMpegCommand(ffc.BuildArgs); err != nil {
		return fmt.Errorf("error within runffmpegcommand(): %s", err)
	}

	return
}

func (ffc *FFConvert) ConvertSingleVideo() (err error) {
	ffc.ConvertVideoPrep()

	if err = ffc.ConvertVideo(); err != nil {
		return fmt.Errorf("error within ConvertVideo(): %s", err)
	}

	return
}

func (ffc *FFConvert) ConvertFolderOfVideos() (err error) {
	var dir []os.DirEntry

	if dir, err = os.ReadDir(ffc.InDir); err != nil {
		return err
	}

	for _, d := range dir {
		if cont := checkDirEntryForContinue(d); cont {
			continue
		}

		if err = ffc.ConvertFolderVideoPrep(d.Name()); err != nil {
			return fmt.Errorf("error within ConvertVideosPrep(): %s", err)
		}

		if err = ffc.ConvertVideo(); err != nil {
			return fmt.Errorf("error within ConvertVideo(): %s", err)
		}
	}

	return
}

func (ffc *FFConvert) FFConvert(inDir, outDir string, preMadeArg int, overwrite bool) (err error) {
	var (
		inAbs string
		info  os.FileInfo
	)

	ffc.newFFConvert(inDir, outDir, preMadeArg, overwrite)

	if inAbs, err = filepath.Abs(ffc.InDir); err != nil {
		err = fmt.Errorf("error getting abs filepath ffconvert(): %w", err)
		return
	}

	if info, err = os.Stat(inAbs); err != nil {
		fmt.Println("os.Stat error: ", err)
		return
	}

	if info.IsDir() {
		ffc.ConvertFolderOfVideos()
	} else {
		ffc.ConvertSingleVideo()
	}

	return nil
}
