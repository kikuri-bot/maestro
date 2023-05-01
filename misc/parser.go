package misc

import (
	"image/color"
	"maestro/config"
	"net/http"
	"strconv"
)

type ViewCardOptions struct {
	ID        uint32
	Dye       color.RGBA
	Glow      bool
	FrameType config.FrameType
}

// Collects card render details from uri params. Will return http error and nil value if any param is missing or invalid.
func ParseCardRenderOptions(w *http.ResponseWriter, r *http.Request, size uint8) []ViewCardOptions {
	res := make([]ViewCardOptions, size)
	var i uint8 = 1

	for i-1 != size {
		iStr := strconv.FormatUint(uint64(i), 10)

		if v := r.URL.Query().Get("id" + iStr); v != "" {
			ID, err := strconv.ParseUint(v, 10, 32)
			if err != nil {
				http.Error(*w, "character ID ("+iStr+") needs to be be uint32", http.StatusUnprocessableEntity)
				return nil
			}
			res[i-1].ID = uint32(ID)
		} else {
			http.Error(*w, "character ID ("+iStr+") uri param is required", http.StatusBadRequest)
			return nil
		}

		if v := r.URL.Query().Get("c" + iStr); v != "" {
			color, err := strconv.ParseUint(v, 10, 32)
			if err != nil {
				http.Error(*w, "dye color value ("+iStr+") needs to be be uint32", http.StatusUnprocessableEntity)
				return nil
			}
			res[i-1].Dye = ColorValueToRGBA(uint32(color))
		} else {
			http.Error(*w, "dye color value ("+iStr+") uri param is required", http.StatusBadRequest)
			return nil
		}

		if v := r.URL.Query().Get("g" + iStr); v != "" {
			g, err := strconv.ParseBool(v)
			if err != nil {
				http.Error(*w, "glow ("+iStr+") needs to be be bool", http.StatusUnprocessableEntity)
				return nil
			}
			res[i-1].Glow = g
		} else {
			http.Error(*w, "glow ("+iStr+") uri param is required", http.StatusBadRequest)
			return nil
		}

		if v := r.URL.Query().Get("ft" + iStr); v != "" {
			fType, err := strconv.ParseUint(v, 10, 8)
			if err != nil {
				http.Error(*w, "frame type ("+iStr+") needs to be be uint8", http.StatusUnprocessableEntity)
				return nil
			}
			res[i-1].FrameType = config.FrameType(fType)
		} else {
			http.Error(*w, "frame type value ("+iStr+") uri param is required", http.StatusBadRequest)
			return nil
		}

		i++
	}

	return res
}
