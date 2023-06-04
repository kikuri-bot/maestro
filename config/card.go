package config

const DEFAULT_DYE_COLOR uint32 = 8289918

const (
	CDN_PATH          = "/etc/kikuri/cdn/"
	CHARACTER_IMAGE_X = 225
	CHARACTER_IMAGE_Y = 350
	CARD_MAX_X        = 335
	CARD_MAX_Y        = 450
)

type FrameType uint8

const (
	DEFAULT_FRAME FrameType = iota
	BETA_FRAME
)

// Max size: 335x450 [px]
// Character image size: 225x350 [px]
var FrameTable = map[FrameType]FrameDetails{
	DEFAULT_FRAME: {
		Name:        "default",
		SizeX:       245,
		SizeY:       370,
		StaticModel: false,
		MaskModel:   true,
	},
	BETA_FRAME: {
		Name:        "beta",
		SizeX:       251,
		SizeY:       376,
		StaticModel: false,
		MaskModel:   true,
	},
}

type FrameDetails struct {
	Name        string
	SizeX       int
	SizeY       int
	StaticModel bool // Static models cannot be dyed.
	MaskModel   bool // Mask models are dyable. Can be combined with static model, making only parts of frame dyable.
}
