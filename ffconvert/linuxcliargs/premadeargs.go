package linuxcliargs

type PreSetArgs int

const (
	Libx264 PreSetArgs = iota
	VTB264
	SVTAV1
	Libx265
)

type PreSet struct {
	ID   PreSetArgs
	Name string
	LCFG LinuxCLICfg
}

func (cfg *LinuxCLICfg) GetPreSets() []PreSet {
	var preSets []PreSet

	preSets = append(preSets, setVcLibX264())
	preSets = append(preSets, setVTB264())
	preSets = append(preSets, setSVTAV1())
	preSets = append(preSets, setLibx265())

	return preSets
}

func setVcLibX264() PreSet {
	newConfig := LinuxCLICfg{
		VCodec:       VcLibx264,
		Mode:         RcCRF,
		CRF:          20,
		Preset:       "veryfast",
		ACodec:       AcAAC,
		AudioBitrate: "160k",
		Container:    Mp4,
		PixFmt:       "yuv420p",
	}

	return PreSet{
		ID:   Libx264,
		Name: "libx264",
		LCFG: newConfig,
	}
}

func setVTB264() PreSet {
	newConfig := LinuxCLICfg{
		VCodec:       VcVTB264,
		Mode:         RcABR,
		VBitrate:     "5000k",
		ACodec:       AcAAC,
		AudioBitrate: "160k",
		Container:    Mp4,
		PixFmt:       "yuv420p",
	}

	return PreSet{
		ID:   VTB264,
		Name: "vtb264",
		LCFG: newConfig,
	}
}

func setSVTAV1() PreSet {
	newConfig := LinuxCLICfg{
		VCodec:       VcSVTAV1,
		Mode:         RcCQ,
		CQ:           30,
		Preset:       "6",
		ACodec:       AcAAC,
		AudioBitrate: "160k",
		Container:    Mp4,
		PixFmt:       "yuv420p",
	}

	return PreSet{
		ID:   SVTAV1,
		Name: "svtav1",
		LCFG: newConfig,
	}
}

func setLibx265() PreSet {
	newConfig := LinuxCLICfg{
		VCodec:       VcLibx265,
		Mode:         RcCRF,
		CRF:          22,
		Preset:       "slow",
		ACodec:       AcAAC,
		AudioBitrate: "160k",
		Container:    Mp4,
		PixFmt:       "yuv420p",
	}

	return PreSet{
		ID:   Libx265,
		Name: "libx265",
		LCFG: newConfig,
	}
}
