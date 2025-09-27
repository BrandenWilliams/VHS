package linuxcliargs

type PremadeArgs int

const (
	Libx264 PremadeArgs = iota
	VTB264
	SVTAV1
	Libx265
)

func (cfg *LinuxCLICfg) SetPreMadeArg(pa PremadeArgs) LinuxCLICfg {
	switch pa {
	case Libx264:
		return cfg.SetVcLibX264()
	case VTB264:
		return cfg.SetVTB264()
	case SVTAV1:
		return cfg.SetSVTAV1()
	case Libx265:
		return cfg.SetLibx265()
	default:
		return cfg.SetLibx265()
	}
}

func (cfg *LinuxCLICfg) SetVcLibX264() LinuxCLICfg {
	return LinuxCLICfg{
		VCodec:       VcLibx264,
		Mode:         RcCRF,
		CRF:          20,
		Preset:       "veryfast",
		ACodec:       AcAAC,
		AudioBitrate: "160k",
		Container:    Mp4,
		PixFmt:       "yuv420p",
	}
}

func (cfg *LinuxCLICfg) SetVTB264() LinuxCLICfg {
	return LinuxCLICfg{
		VCodec:       VcVTB264,
		Mode:         RcABR,
		VBitrate:     "5000k",
		ACodec:       AcAAC,
		AudioBitrate: "160k",
		Container:    Mp4,
		PixFmt:       "yuv420p",
	}
}

func (cfg *LinuxCLICfg) SetSVTAV1() LinuxCLICfg {
	return LinuxCLICfg{
		VCodec:       VcSVTAV1,
		Mode:         RcCQ,
		CQ:           30,
		Preset:       "6",
		ACodec:       AcAAC,
		AudioBitrate: "160k",
		Container:    Mp4,
		PixFmt:       "yuv420p",
	}
}

func (cfg *LinuxCLICfg) SetLibx265() LinuxCLICfg {
	return LinuxCLICfg{
		VCodec:       VcLibx265,
		Mode:         RcCRF,
		CRF:          22,
		Preset:       "slow",
		ACodec:       AcAAC,
		AudioBitrate: "160k",
		Container:    Mp4,
		PixFmt:       "yuv420p",
	}
}
