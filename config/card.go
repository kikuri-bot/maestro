package config

const (
	CHARACTER_IMAGE_X = 225
	CHARACTER_IMAGE_Y = 350
	CARD_MAX_X        = 335
	CARD_MAX_Y        = 450
)

type FrameType uint8

const (
	DEFAULT_FRAME FrameType = iota
	BETA_FRAME
	BETA2_FRAME
)

// Max size: 335x450 [px]
// Character image size: 225x350 [px]
var FrameTable = map[FrameType]FrameDetails{
	DEFAULT_FRAME: {
		Name:          "default",
		SizeX:         245,
		SizeY:         370,
		TwoLayerModel: false, // If true, it means there's 2 parts of frame - static and editable one.
		Dyeable:       true,
	},
	BETA_FRAME: {
		Name:          "beta",
		SizeX:         251,
		SizeY:         376,
		TwoLayerModel: false,
		Dyeable:       true,
	},
}

type FrameDetails struct {
	Name          string
	SizeX         int
	SizeY         int
	TwoLayerModel bool
	Dyeable       bool
}
