package main

import (
	"flag"
	"maestro/logger"
	routeCard "maestro/route/card"
	"net/http"
)

func main() {
	logger.InitLogger()
	logger.Info.Println("Starting Maestro application...")

	// http.HandleFunc("/card/album", routeCard.AlbumHandler)
	// http.HandleFunc("/card/dye", routeCard.DyeHandler)
	// http.HandleFunc("/card/preview", routeCard.PreviewHandler)
	http.HandleFunc("/card/view", routeCard.ViewHandler)

	var addr string
	flag.StringVar(&addr, "addr", "127.0.0.1:8808", "switches which address should be used")
	flag.Parse()

	logger.Info.Printf("Started application on %s.\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		logger.Error.Panic(err)
	}
}
