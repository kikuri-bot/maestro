package route

import (
	"image"
	"image/png"
	"maestro/config"
	"maestro/misc"
	"maestro/render"
	"net/http"
	"strconv"
)

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var (
		opt render.CardRenderOptions
		fd  config.FrameDetails
	)

	if v := r.URL.Query().Get("id"); v != "" {
		ID, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			http.Error(w, "character ID needs to be be uint32", http.StatusUnprocessableEntity)
			return
		}
		opt.CharacterID = uint32(ID)
	} else {
		http.Error(w, "character ID uri param is required", http.StatusBadRequest)
		return
	}

	if v := r.URL.Query().Get("c"); v != "" {
		color, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			http.Error(w, "dye color value needs to be be uint32", http.StatusUnprocessableEntity)
			return
		}
		opt.Dye = misc.ColorValueToRGBA(uint32(color))
	} else {
		http.Error(w, "dye color value uri param is required", http.StatusBadRequest)
		return
	}

	if v := r.URL.Query().Get("g"); v != "" {
		g, err := strconv.ParseBool(v)
		if err != nil {
			http.Error(w, "glow needs to be be bool", http.StatusUnprocessableEntity)
			return
		}
		opt.Glow = g
	} else {
		http.Error(w, "glow uri param is required", http.StatusBadRequest)
		return
	}

	if v := r.URL.Query().Get("ft"); v != "" {
		fType, err := strconv.ParseUint(v, 10, 8)
		if err != nil {
			http.Error(w, "frame type needs to be be uint8", http.StatusUnprocessableEntity)
			return
		}
		ft := config.FrameType(fType)
		fd = config.FrameTable[ft]
		opt.FrameType = ft
	} else {
		http.Error(w, "frame type uri param is required", http.StatusBadRequest)
		return
	}

	canvas := image.NewRGBA(image.Rectangle{
		image.Point{},
		image.Point{
			fd.SizeX,
			fd.SizeY,
		},
	})

	canvas, err := render.DrawCard(canvas, opt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	err = png.Encode(w, canvas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
