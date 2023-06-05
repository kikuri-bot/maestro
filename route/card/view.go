package route

import (
	"image/png"
	"maestro/logger"
	"maestro/misc"
	"maestro/render"
	"net/http"
	"time"
)

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	startedAt := time.Now()
	cardRenderOptions, err := misc.ParseCardRenderOptions(r.URL.Query(), 1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	canvas, err := render.DrawCard(nil, cardRenderOptions[0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	took := time.Since(startedAt)
	if took > time.Second*3 {
		logger.Warn.Printf("Took %s for single card render request!\n", took.String())
	}

	w.Header().Set("X-Processing-Time", took.String())
	w.Header().Set("Content-Type", "image/png")
	enc := &png.Encoder{
		CompressionLevel: png.NoCompression,
	}

	if err = enc.Encode(w, canvas); err != nil {
		logger.Warn.Println("Failed to encode finished card render to response: ", err.Error())
	}
}
