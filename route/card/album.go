package route

import (
	"image"
	"image/png"
	"maestro/config"
	"maestro/misc"
	"maestro/render"
	"math"
	"net/http"
	"strconv"
)

func AlbumHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var amount uint8
	if v := r.URL.Query().Get("q"); v != "" {
		quantity, err := strconv.ParseUint(v, 10, 8)
		if err != nil {
			http.Error(w, "quantity needs to be be uint8", http.StatusUnprocessableEntity)
			return
		}

		if quantity > 10 {
			http.Error(w, "quantity cannot be greater than 10", http.StatusUnprocessableEntity)
			return
		} else if quantity == 0 {
			http.Error(w, "quantity needs to be 1 or greater", http.StatusUnprocessableEntity)
			return
		}

		amount = uint8(quantity)
	} else {
		http.Error(w, "quantity uri param is required", http.StatusBadRequest)
		return
	}

	var columns int
	if v := r.URL.Query().Get("cl"); v != "" {
		value, err := strconv.ParseUint(v, 10, 8)
		if err != nil {
			http.Error(w, "column number needs to be be uint8", http.StatusUnprocessableEntity)
			return
		}

		if value > 5 {
			http.Error(w, "column number cannot be greater than 5", http.StatusUnprocessableEntity)
			return
		} else if value == 0 {
			http.Error(w, "column number needs to be 1 or greater", http.StatusUnprocessableEntity)
			return
		}

		columns = int(value)
	} else {
		http.Error(w, "column number uri param is required", http.StatusBadRequest)
		return
	}

	opts := misc.ParseCardRenderOptions(&w, r, amount)
	if opts == nil {
		return
	}

	rows := int(math.Ceil(float64(amount) / float64(columns)))
	if rows*columns < int(amount) {
		http.Error(w, "provided quantity is larger than declared grid", http.StatusBadRequest)
		return
	}

	canvas := image.NewRGBA(image.Rectangle{
		image.Point{},
		image.Point{
			config.CARD_MAX_X*columns + (8 * columns),
			config.CARD_MAX_Y*rows + (8 * rows),
		},
	})

	var err error
	for idx, opt := range opts {
		y := int(math.Floor(float64(idx) / float64(columns)))
		x := idx - y*columns

		frameDetails := config.FrameTable[opt.FrameType]
		canvas, err = render.DrawCard(canvas, render.CardRenderOptions{
			CharacterID: opt.ID,
			FrameType:   opt.FrameType,
			Dye:         opt.Dye,
			Glow:        opt.Glow,
			OffsetX:     x*config.CARD_MAX_X + (8 * x) + (config.CARD_MAX_X-frameDetails.SizeX)/2,
			OffsetY:     y*config.CARD_MAX_Y + (8 * y) + (config.CARD_MAX_Y-frameDetails.SizeY)/2,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "image/png")
	err = png.Encode(w, canvas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
