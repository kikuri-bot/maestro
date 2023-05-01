package render

import (
	"errors"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"maestro/misc"
	"os"
)

type DyeEmbedRenderOptions struct {
	Color   color.RGBA
	Glow    bool
	OffsetX int
	OffsetY int
}

func DrawDyeEmbed(canvas *image.RGBA, opt DyeEmbedRenderOptions) (*image.RGBA, error) {
	var (
		baseImgBuf *os.File
		err        error
	)

	if opt.Glow {
		baseImgBuf, err = os.OpenFile("../cdn/private/dye-embed/glow-base.png", os.O_RDONLY, 0755)
	} else {
		baseImgBuf, err = os.OpenFile("../cdn/private/dye-embed/base.png", os.O_RDONLY, 0755)
	}

	if err != nil {
		return nil, errors.New("requested dye (base model) doesn't exist")
	}

	baseImage, err := png.Decode(baseImgBuf)
	if err != nil {
		return nil, errors.New("requested dye (base model) image file is malformed:" + err.Error())
	}
	defer baseImgBuf.Close()

	dimensions := image.Rectangle{
		image.Point{
			opt.OffsetX,
			opt.OffsetY,
		},
		image.Point{
			opt.OffsetX + 512,
			opt.OffsetY + 512,
		},
	}

	draw.Draw(
		canvas,
		dimensions,
		baseImage,
		image.Point{},
		draw.Src,
	)

	var maskImgBuf *os.File
	if opt.Glow {
		maskImgBuf, err = os.OpenFile("../cdn/private/dye-embed/glow-mask.png", os.O_RDONLY, 0755)
	} else {
		maskImgBuf, err = os.OpenFile("../cdn/private/dye-embed/mask.png", os.O_RDONLY, 0755)
	}

	if err != nil {
		return nil, errors.New("requested dye (mask model) doesn't exist")
	}

	maskImage, err := png.Decode(maskImgBuf)
	if err != nil {
		return nil, errors.New("requested dye (mask model) image file is malformed:" + err.Error())
	}
	defer maskImgBuf.Close()

	draw.Draw(
		canvas,
		dimensions,
		misc.Recolor(maskImage, opt.Color),
		image.Point{},
		draw.Over,
	)

	return canvas, nil
}
