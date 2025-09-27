package linuxcliargs

import (
	"strconv"
	"strings"
)

func getExtraInputs(eia []string) (args []string) {
	if len(eia) > 0 {
		args = append(args, eia...)
	}

	return args
}

func getVCArg(vc VideoCodec) (args []string) {
	if string(vc) == "" {
		args = append(args, "-c:v")
		args = append(args, string(VcLibx264))
		return
	}

	args = append(args, "-c:v")
	args = append(args, string(vc))
	return
}

func getFilters(scale string, extraFilterArgs []string) (fstr []string) {
	// filters
	var filters []string
	if scale != "" {
		filters = append(filters, "scale="+scale)
	}

	if len(extraFilterArgs) > 0 {
		filters = append(filters, extraFilterArgs...)
	}

	if len(filters) > 0 {
		fstr = append(fstr, "-vf", strings.Join(filters, ","))
		return
	}

	return
}

func rateControlPerCodec(lca LinuxCLICfg) (args []string) {
	switch lca.Mode {
	case RcCRF:
		if lca.CRF == 0 {
			lca.CRF = 20
		}

		args = append(args, "-crf", strconv.Itoa(lca.CRF))
		if lca.Preset != "" {
			args = append(args, "-preset", lca.Preset)
		}
	case RcCQ:
		cq := lca.CQ
		if cq == 0 {
			cq = 30
		}

		args = append(args, "-crf", strconv.Itoa(cq))
		if lca.Preset != "" {
			args = append(args, "-preset", lca.Preset)
		}
	case RcCBR, RcABR:
		if lca.VBitrate != "" {
			args = append(args, "-b:v", lca.VBitrate)
		}
		if lca.Maxrate != "" {
			args = append(args, "-maxrate", lca.Maxrate)
		}
		if lca.Bufsize != "" {
			args = append(args, "-bufsize", lca.Bufsize)
		}
		if lca.Preset != "" {
			args = append(args, "-preset", lca.Preset)
		}
	default:
		args = append(args, "-crf", "20", "-preset", "veryfast")
	}

	return
}

func getTuneArgs(tune string) (args []string) {
	if tune != "" {
		args = append(args, "-tune")
		args = append(args, tune)

		return
	}

	return
}

func getProfileArgs(profile string) (args []string) {
	if profile != "" {
		args = append(args, "-profile:v")
		args = append(args, profile)
		return
	}

	return
}

func getLevelArgs(level string) (args []string) {
	if level != "" {
		args = append(args, "-level:v")
		args = append(args, level)
		return
	}

	return
}

func getPixFMTArgs(fixfmt string) (args []string) {
	if fixfmt != "" {
		args = append(args, "-pix_fmt", fixfmt)
		return
	}

	return
}

func getFPSArgs(fps string) (args []string) {
	if fps != "" {
		args = append(args, "-r")
		args = append(args, fps)
		return
	}

	return
}

func getAudioArgs(ac AudioCodec, bitrate string, ACh int, ARate int) (args []string) {
	if ac == "" {
		ac = AcAAC
	}

	args = append(args, "-c:a", string(ac))
	if bitrate != "" {
		args = append(args, "-b:a", bitrate)
	}
	if ACh > 0 {
		args = append(args, "-ac", strconv.Itoa(ACh))
	}
	if ARate > 0 {
		args = append(args, "-ar", strconv.Itoa(ARate))
	}

	return
}

func getContainerArgs(c Container) (args []string) {
	if c == Mp4 {
		args = append(args, "-movflags")
		args = append(args, "+faststart")
	}

	return
}

func getExtraOutputArgs(eoa []string) (args []string) {
	if len(eoa) > 0 {
		args = append(args, eoa...)
		return
	}

	return
}
