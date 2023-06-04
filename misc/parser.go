package misc

import (
	"fmt"
	"image/color"
	"maestro/config"
	"net/url"
	"strconv"
)

type CardRender struct {
	ID        uint32
	Dye       color.RGBA
	Glow      bool
	FrameType config.FrameType
}

// Collects card render details from uri params. Will return http error and nil value if any param is missing or invalid.
func ParseCardRenderOptions(uriValues url.Values, size uint8) ([]CardRender, error) {
	res := make([]CardRender, size)
	var i uint8 = 0

	for i != size {
		iStr := strconv.FormatUint(uint64(i), 10)

		if v := uriValues.Get("id" + iStr); v != "" {
			ID, err := strconv.ParseUint(v, 10, 32)
			if err != nil {
				return nil, fmt.Errorf("CardRender[%s].id must be of type uint32", iStr)
			}

			res[i].ID = uint32(ID)
		} else {
			return nil, fmt.Errorf("CardRender[%s].id is required", iStr)
		}

		if v := uriValues.Get("c" + iStr); v != "" {
			c, err := strconv.ParseUint(v, 10, 32)
			if err != nil {
				return nil, fmt.Errorf("CardRender[%s].c must be of type uint32", iStr)
			}

			res[i].Dye = ColorValueToRGBA(uint32(c))
		} else {
			res[i].Dye = ColorValueToRGBA(config.DEFAULT_DYE_COLOR)
		}

		if v := uriValues.Get("g" + iStr); v != "" {
			g, err := strconv.ParseBool(v)
			if err != nil {
				return nil, fmt.Errorf("CardRender[%s].g must be of type bool", iStr)
			}

			res[i].Glow = g
		}

		if v := uriValues.Get("f" + iStr); v != "" {
			fType, err := strconv.ParseUint(v, 10, 8)
			if err != nil {
				return nil, fmt.Errorf("CardRender[%s].f must be of type uint8", iStr)
			}

			res[i].FrameType = config.FrameType(fType)
		}

		i++
	}

	return res, nil
}
