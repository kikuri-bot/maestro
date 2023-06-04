package route

// import (
// 	"image"
// 	"image/png"
// 	"maestro/misc"
// 	"maestro/render"
// 	"net/http"
// 	"strconv"
// )

// func DyeHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodGet {
// 		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	var opt render.DyeEmbedRenderOptions
// 	if v := r.URL.Query().Get("c"); v != "" {
// 		color, err := strconv.ParseUint(v, 10, 32)
// 		if err != nil {
// 			http.Error(w, "dye color value needs to be be uint32", http.StatusUnprocessableEntity)
// 			return
// 		}
// 		opt.Color = misc.ColorValueToRGBA(uint32(color))
// 	} else {
// 		http.Error(w, "dye color value uri param is required", http.StatusBadRequest)
// 		return
// 	}

// 	if v := r.URL.Query().Get("g"); v != "" {
// 		g, err := strconv.ParseBool(v)
// 		if err != nil {
// 			http.Error(w, "glow needs to be be bool", http.StatusUnprocessableEntity)
// 			return
// 		}
// 		opt.Glow = g
// 	} else {
// 		http.Error(w, "glow uri param is required", http.StatusBadRequest)
// 		return
// 	}

// 	canvas := image.NewRGBA(image.Rectangle{
// 		image.Point{},
// 		image.Point{
// 			512,
// 			512,
// 		},
// 	})

// 	canvas, err := render.DrawDyeEmbed(canvas, opt)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "image/png")
// 	err = png.Encode(w, canvas)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }
