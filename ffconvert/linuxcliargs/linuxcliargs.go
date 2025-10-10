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
	VCodec    VideoCodec `json:"vcodec" form:"vcodec"`
	ACodec    AudioCodec `json:"acodec" form:"acodec"`
	Container Container  `json:"container" form:"container"`

	// rate control
	Mode RateControlMode `json:"mode" form:"mode"`

	CRF      int    `json:"crf" form:"crf"`
	CQ       int    `json:"cq" form:"cq"`
	VBitrate string `json:"vbitrate" form:"vbitrate"`
	Maxrate  string `json:"maxrate" form:"maxrate"`
	Bufsize  string `json:"buffsize" form:"buffsize"`

	// x264 x265 specific args
	Preset  string `json:"preset" form:"preset"`
	Tune    string `json:"tune" form:"tune"`
	Profile string `json:"profile" form:"profile"`
	Level   string `json:"level" form:"level"`

	//extras
	PixFmt       string `json:"pixfmt" form:"pixfmt"`
	Scale        string `json:"scale" form:"scale"`
	Fps          string `json:"fps" form:"fps"`
	AudioBitrate string `json:"audiobitrate" form:"audiobitrate"`
	AudioCh      int    `json:"audoch" form:"audoch"`
	AudioRate    int    `json:"audiorate" form:"audiorate"`

	ExtraInputArgs  []string `json:"extrainputargs" form:"extrainputargs"`
	ExtraFilterArgs []string `json:"extrafilterargs" form:"extrafilterargs"`
	ExtraOutputArgs []string `json:"extraoutputargs" form:"extraoutputargs"`
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
