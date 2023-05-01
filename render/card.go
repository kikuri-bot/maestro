package render

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"maestro/config"
	"maestro/misc"
	"os"
)

type CardRenderOptions struct {
	CharacterID uint32
	FrameType   config.FrameType
	Dye         color.RGBA
	Glow        bool
	OffsetX     int
	OffsetY     int
}

func DrawCard(canvas *image.RGBA, opt CardRenderOptions) (*image.RGBA, error) {
	charImgBuf, err := os.OpenFile(fmt.Sprintf("../cdn/public/character/%d.png", opt.CharacterID), os.O_RDONLY, 0755)
	if err != nil {
		return nil, errors.New("requested character doesn't exist")
	}

	characterImage, err := png.Decode(charImgBuf)
	if err != nil {
		return nil, errors.New("requested character's image file is malformed:" + err.Error())
	}
	defer charImgBuf.Close()

	frameDetails := config.FrameTable[opt.FrameType]
	charImgOffsetX := opt.OffsetX + ((frameDetails.SizeX - config.CHARACTER_IMAGE_X) / 2)
	charImgOffsetY := opt.OffsetY + ((frameDetails.SizeY - config.CHARACTER_IMAGE_Y) / 2)

	draw.Draw(
		canvas,
		image.Rectangle{
			image.Point{
				charImgOffsetX,
				charImgOffsetY,
			},
			image.Point{
				config.CHARACTER_IMAGE_X + charImgOffsetX,
				config.CHARACTER_IMAGE_Y + charImgOffsetY,
			},
		},
		characterImage,
		image.Point{},
		draw.Src,
	)

	// Draw static model of frame on character's image.
	if frameDetails.TwoLayerModel || !frameDetails.Dyeable {
		staticFrameBuf, err := os.OpenFile("../cdn/private/frame/"+frameDetails.Name+"/base.png", os.O_RDONLY, 0755)
		if err != nil {
			return nil, errors.New("requested frame (base model) doesn't exist")
		}

		staticFrameImage, err := png.Decode(staticFrameBuf)
		if err != nil {
			return nil, errors.New("requested frame (base model) image file is malformed:" + err.Error())
		}
		defer staticFrameBuf.Close()

		draw.Draw(
			canvas,
			image.Rectangle{
				image.Point{opt.OffsetX, opt.OffsetY},
				image.Point{frameDetails.SizeX + opt.OffsetX, frameDetails.SizeY + opt.OffsetY},
			},
			staticFrameImage,
			image.Point{},
			draw.Over,
		)
	}

	if frameDetails.Dyeable && (opt.Dye.R != 0 || opt.Dye.G != 0 || opt.Dye.B != 0 || opt.Dye.A != 0) {
		var (
			dyeableFrameBuf *os.File
			err             error
		)

		if opt.Glow {
			dyeableFrameBuf, err = os.OpenFile("../cdn/private/frame/"+frameDetails.Name+"/glow-mask.png", os.O_RDONLY, 0755)
		} else {
			dyeableFrameBuf, err = os.OpenFile("../cdn/private/frame/"+frameDetails.Name+"/mask.png", os.O_RDONLY, 0755)
		}

		if err != nil {
			return nil, errors.New("requested frame (mask model) doesn't exist")
		}

		dyeableFrameImage, err := png.Decode(dyeableFrameBuf)
		if err != nil {
			return nil, errors.New("requested frame (mask model) image file is malformed:" + err.Error())
		}
		defer dyeableFrameBuf.Close()

		draw.Draw(
			canvas,
			image.Rectangle{
				image.Point{opt.OffsetX, opt.OffsetY},
				image.Point{frameDetails.SizeX + opt.OffsetX, frameDetails.SizeY + opt.OffsetY},
			},
			misc.Recolor(dyeableFrameImage, opt.Dye),
			image.Point{},
			draw.Over,
		)
	}

	return canvas, nil
}
