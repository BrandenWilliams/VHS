package ffconvert

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/BrandenWilliams/VHS/ffconvert/linuxcliargs"
)

type FFConvert struct {
	PreSets   []linuxcliargs.PreSet
	LCliA     linuxcliargs.LinuxCLICfg
	BuildArgs []string

	InDir  string
	OutDir string

	InAbsDir  string
	OutAbsDir string

	CurrentFile string

	Crf       int
	Preset    string
	Overwrite bool
}

func (ffc *FFConvert) SetPreSets() {
	ffc.PreSets = ffc.LCliA.GetPreSets()
}

func (ffc *FFConvert) SetLinuxCLIConfig(preMadeArg linuxcliargs.PreSetArgs) (err error) {
	for _, set := range ffc.PreSets {
		if set.ID == preMadeArg {
			ffc.LCliA = set.LCFG
		}
	}

	return fmt.Errorf("linux cli config preset not found")
}

func (ffc *FFConvert) NewFFConvert(inDir, outDir string, crf int, preset string, overwrite bool, preMadeArg linuxcliargs.PreSetArgs) (err error) {
	ffc.InDir = inDir
	ffc.OutDir = outDir
	ffc.Crf = crf
	ffc.Preset = preset
	ffc.Overwrite = overwrite

	if err = ffc.SetLinuxCLIConfig(preMadeArg); err != nil {
		return err
	}

	return
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
			return fmt.Errorf("error within ConvertVideo(): %s", err)
		}
	}

	return nil
}

func (ffc *FFConvert) RunFFMpegCommand() (err error) {
	if len(ffc.BuildArgs) == 0 {
		return fmt.Errorf("empty args string failed: %s", err)
	}

	cmd := exec.Command("ffmpeg", ffc.BuildArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err = cmd.Run(); err != nil {
		return fmt.Errorf("cmdrun() ffmpeg failed: %w", err)
	}

	return nil
}

func (ffc *FFConvert) ConvertVideo(currentFile string) (err error) {
	if ffc.InAbsDir, ffc.OutAbsDir, err = getAllAbsolutePaths(ffc.InDir+currentFile, ffc.OutDir); err != nil {
		return fmt.Errorf("error within getallabsolutepaths(): %s", ffc.InDir)
	}

	if ffc.OutAbsDir, err = generateOutPathArg(currentFile, ffc.OutAbsDir, ".mp4"); err != nil {
		return fmt.Errorf("error within generateoutpatharg(): %s", err)
	}

	// LONGTERM NEED TO ADD OUTPUTFILE COUNTER FOR SAME FOLDER SUPPORT
	if checkIfSameFilePath(ffc.InAbsDir, ffc.OutAbsDir) {
		return fmt.Errorf("input file and output file resolve to the same file: %s", ffc.InAbsDir)
	}

	if ffc.BuildArgs, err = ffc.LCliA.BuildLinuxCLIArgs(ffc.InAbsDir, ffc.OutAbsDir, ffc.Overwrite); err != nil {
		return fmt.Errorf("error within buildlinuxcliargs(): %s", err)
	}

	PrintLinuxCLIArgs(ffc.BuildArgs)

	if err = ffc.RunFFMpegCommand(); err != nil {
		return fmt.Errorf("error within runffmpegcommand(): %s", err)
	}

	return nil
}

// DEBUG FUNCTIONS
func PrintLinuxCLIArgs(bArgs []string) {
	log.Printf("LinuxCLIArgs() print start\n")
	var newBArgs string

	for _, s := range bArgs {
		newBArgs += s
	}

	log.Printf("newBArgs: %v", bArgs)
	log.Printf("\nLinuxCLIArgs() print stop\n")
}
