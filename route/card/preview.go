package route

// func PreviewHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodGet {
// 		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	opts := misc.ParseCardRenderOptions(&w, r, 2)
// 	if opts == nil {
// 		return
// 	}

// 	canvas := image.NewRGBA(image.Rectangle{
// 		image.Point{},
// 		image.Point{
// 			config.CARD_MAX_X*2 + 128,
// 			config.CARD_MAX_Y,
// 		},
// 	})

// 	f1 := config.FrameTable[opts[0].FrameType]
// 	canvas, err := render.DrawCard(canvas, render.CardRenderOptions{
// 		CharacterID: opts[0].ID,
// 		FrameType:   opts[0].FrameType,
// 		Dye:         opts[0].Dye,
// 		Glow:        opts[0].Glow,
// 		OffsetX:     (config.CARD_MAX_X - f1.SizeX) / 2,
// 		OffsetY:     (config.CARD_MAX_Y - f1.SizeY) / 2,
// 	})

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	arrowImgBuf, err := os.OpenFile("../cdn/private/render-misc/multi-preview-arrow.png", os.O_RDONLY, 0755)
// 	if err != nil {
// 		http.Error(w, "arrow icon image doesn't exist on disk", http.StatusInternalServerError)
// 		return
// 	}

// 	arrowImage, err := png.Decode(arrowImgBuf)
// 	if err != nil {
// 		http.Error(w, "arrow icon image is malformed:"+err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	defer arrowImgBuf.Close()

// 	draw.Draw(
// 		canvas,
// 		// image.Rectangle{
// 		// 	image.Point{
// 		// 		config.CARD_MAX_X + 8,
// 		// 		200,
// 		// 	},
// 		// 	image.Point{
// 		// 		config.CARD_MAX_X + 120,
// 		// 		249,
// 		// 	},
// 		// },
// 		image.Rectangle{
// 			image.Point{
// 				config.CARD_MAX_X + 8,
// 				95,
// 			},
// 			image.Point{
// 				config.CARD_MAX_X + 120,
// 				355,
// 			},
// 		},
// 		arrowImage,
// 		image.Point{},
// 		draw.Src,
// 	)

// 	f2 := config.FrameTable[opts[1].FrameType]
// 	canvas, err = render.DrawCard(canvas, render.CardRenderOptions{
// 		CharacterID: opts[1].ID,
// 		FrameType:   opts[1].FrameType,
// 		Dye:         opts[1].Dye,
// 		Glow:        opts[1].Glow,
// 		OffsetX:     config.CARD_MAX_X + 128 + (config.CARD_MAX_X-f2.SizeX)/2,
// 		OffsetY:     (config.CARD_MAX_Y - f2.SizeY) / 2,
// 	})

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
