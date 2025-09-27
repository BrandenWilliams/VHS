package linuxcliargs

import "fmt"

type (
	VideoCodec      string
	RateControlMode string
	AudioCodec      string
	Container       string
)

const (
	// VideoCodecs
	VcLibx264  VideoCodec = "libx264"
	VcVTB264   VideoCodec = "h264_videotoolbox"
	VcNVENC264 VideoCodec = "h264_nvenc"
	VcLibx265  VideoCodec = "libx265"
	VcVTB265   VideoCodec = "hevc_videotoolbox"
	VcNVENC265 VideoCodec = "hevc_nvenc"
	VcSVTAV1   VideoCodec = "libsvtav1"
	VcAOMAV1   VideoCodec = "libaom-av1"
	VcVP9      VideoCodec = "libvpx-vp9"

	// AudioCodec
	AcAAC  AudioCodec = "aac"
	AcOpus AudioCodec = "libopus"
	AcAC3  AudioCodec = "ac3"

	// Container
	Mp4  Container = "mp4"
	Mkv  Container = "mkv"
	WebM Container = "webm"

	// RateControlModes
	RcCRF RateControlMode = "crf"
	RcCQ  RateControlMode = "cq"
	RcCBR RateControlMode = "cbr"
	RcABR RateControlMode = "abr"
)

type LinuxCLICfg struct {
	VCodec    VideoCodec
	ACodec    AudioCodec
	Container Container

	// rate control
	Mode     RateControlMode
	CRF      int
	CQ       int
	VBitrate string
	Maxrate  string
	Bufsize  string

	// x264 x265 specific args
	Preset  string
	Tune    string
	Profile string
	Level   string

	//extras
	PixFmt       string
	Scale        string
	Fps          string
	AudioBitrate string
	AudioCh      int
	AudioRate    int

	ExtraInputArgs  []string
	ExtraFilterArgs []string
	ExtraOutputArgs []string
}

func (lca *LinuxCLICfg) BuildLinuxCLIArgs(inPath, outPath string, overwrite bool) (args []string, err error) {
	if overwrite {
		args = []string{"-y"}
	}

	args = append(args, getExtraInputs(lca.ExtraInputArgs)...)

	args = append(args, "-i", inPath)

	args = append(args, getFilters(lca.Scale, lca.ExtraFilterArgs)...)

	args = append(args, getVCArg(lca.VCodec)...)

	args = append(args, rateControlPerCodec(*lca)...)

	args = append(args, getTuneArgs(lca.Tune)...)

	args = append(args, getProfileArgs(lca.Profile)...)

	args = append(args, getLevelArgs(lca.Level)...)

	args = append(args, getPixFMTArgs(lca.PixFmt)...)

	args = append(args, getFPSArgs(lca.Fps)...)

	args = append(args, getAudioArgs(lca.ACodec, lca.AudioBitrate, lca.AudioCh, lca.AudioRate)...)

	args = append(args, getContainerArgs(lca.Container)...)

	args = append(args, getExtraOutputArgs(lca.ExtraOutputArgs)...)

	args = append(args, outPath)

	if len(args) == 0 {
		return args, fmt.Errorf("linux cli argument empty")
	}

	return
}
