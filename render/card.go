package render

import (
	"context"
	"errors"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"maestro/config"
	"maestro/logger"
	"maestro/misc"
	"os"
)

func DrawCard(ctx context.Context, canvas *image.RGBA, opt misc.CardRenderOptions) (*image.RGBA, error) {
	if int(opt.FrameType) > len(config.FrameTable) {
		return nil, errors.New("requested frame doesn't exist")
	}

	charImgPath := fmt.Sprintf("%s/public/character/%d.png", config.CDN_PATH, opt.ID)
	charImgBuf, err := os.OpenFile(charImgPath, os.O_RDONLY, 0755)
	if err != nil {
		return nil, fmt.Errorf("requested character (ID = %d) doesn't exist", opt.ID)
	}

	if ctx.Err() != nil {
		logger.Warn.Println("Failed to render time within allowed context, original error: ", ctx.Err().Error())
		return nil, errors.New("failed to render card within allowed context - please try again later and report this issue if it continues to happen")
	}

	characterImage, err := png.Decode(charImgBuf)
	if err != nil {
		logger.Error.Printf(`Failed to decode image of "%s" character. Original error: %s\n`, charImgPath, err.Error())
		return nil, errors.New("failed to render character's card - please try again later and report this issue if it continues to happen")
	}
	defer charImgBuf.Close()

	if ctx.Err() != nil {
		logger.Warn.Println("Failed to render time within allowed context, original error: ", ctx.Err().Error())
		return nil, errors.New("failed to render card within allowed context - please try again later and report this issue if it continues to happen")
	}

	frameDetails := config.FrameTable[opt.FrameType]
	charImgOffsetX := opt.OffsetX + ((frameDetails.SizeX - config.CHARACTER_IMAGE_X) / 2)
	charImgOffsetY := opt.OffsetY + ((frameDetails.SizeY - config.CHARACTER_IMAGE_Y) / 2)

	if canvas == nil {
		canvas = image.NewRGBA(image.Rectangle{
			image.Point{},
			image.Point{
				frameDetails.SizeX,
				frameDetails.SizeY,
			},
		})
	}

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

	if ctx.Err() != nil {
		logger.Warn.Println("Failed to render time within allowed context, original error: ", ctx.Err().Error())
		return nil, errors.New("failed to render card within allowed context - please try again later and report this issue if it continues to happen")
	}

	if frameDetails.MaskModel {
		var (
			maskFrameBuf *os.File
			err          error
		)

		if opt.Glow {
			maskFrameBuf, err = os.OpenFile("../cdn/private/frame/"+frameDetails.Name+"/glow-mask.png", os.O_RDONLY, 0755)
		} else {
			maskFrameBuf, err = os.OpenFile("../cdn/private/frame/"+frameDetails.Name+"/mask.png", os.O_RDONLY, 0755)
		}

		if err != nil {
			return nil, errors.New("requested frame (mask model) doesn't exist")
		}

		if ctx.Err() != nil {
			logger.Warn.Println("Failed to render time within allowed context, original error: ", ctx.Err().Error())
			return nil, errors.New("failed to render card within allowed context - please try again later and report this issue if it continues to happen")
		}

		maskFrameImage, err := png.Decode(maskFrameBuf)
		if err != nil {
			logger.Error.Printf(`Failed to decode "%s" frame (ID = %d) mask model. Original error: %s\n`, frameDetails.Name, opt.FrameType, err.Error())
			return nil, errors.New("failed to render character's card - please try again later and report this issue if it continues to happen")
		}
		defer maskFrameBuf.Close()

		if ctx.Err() != nil {
			logger.Warn.Println("Failed to render time within allowed context, original error: ", ctx.Err().Error())
			return nil, errors.New("failed to render card within allowed context - please try again later and report this issue if it continues to happen")
		}

		draw.Draw(
			canvas,
			image.Rectangle{
				image.Point{opt.OffsetX, opt.OffsetY},
				image.Point{frameDetails.SizeX + opt.OffsetX, frameDetails.SizeY + opt.OffsetY},
			},
			misc.Recolor(maskFrameImage, opt.Dye),
			image.Point{},
			draw.Over,
		)

		if ctx.Err() != nil {
			logger.Warn.Println("Failed to render time within allowed context, original error: ", ctx.Err().Error())
			return nil, errors.New("failed to render card within allowed context - please try again later and report this issue if it continues to happen")
		}
	}

	if frameDetails.StaticModel {
		var (
			staticFrameBuf *os.File
			err            error
		)

		if opt.Glow {
			staticFrameBuf, err = os.OpenFile("../cdn/private/frame/"+frameDetails.Name+"/glow-static.png", os.O_RDONLY, 0755)
		} else {
			staticFrameBuf, err = os.OpenFile("../cdn/private/frame/"+frameDetails.Name+"/static.png", os.O_RDONLY, 0755)
		}

		if err != nil {
			return nil, errors.New("requested frame (static model) doesn't exist")
		}

		if ctx.Err() != nil {
			logger.Warn.Println("Failed to render time within allowed context, original error: ", ctx.Err().Error())
			return nil, errors.New("failed to render card within allowed context - please try again later and report this issue if it continues to happen")
		}

		staticFrameImage, err := png.Decode(staticFrameBuf)
		if err != nil {
			logger.Error.Printf(`Failed to decode "%s" frame (ID = %d) static model. Original error: %s\n`, frameDetails.Name, opt.FrameType, err.Error())
			return nil, errors.New("failed to render character's card - please try again later and report this issue if it continues to happen")
		}
		defer staticFrameBuf.Close()

		if ctx.Err() != nil {
			logger.Warn.Println("Failed to render time within allowed context, original error: ", ctx.Err().Error())
			return nil, errors.New("failed to render card within allowed context - please try again later and report this issue if it continues to happen")
		}

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

		if ctx.Err() != nil {
			logger.Warn.Println("Failed to render time within allowed context, original error: ", ctx.Err().Error())
			return nil, errors.New("failed to render card within allowed context - please try again later and report this issue if it continues to happen")
		}
	}

	return canvas, nil
}
