package main

import (
	"flag"
	"maestro/logger"
	routeCard "maestro/route/card"
	"net/http"
	"os"

	godotenv "github.com/joho/godotenv"
)

func main() {
	logger.InitLogger()
	logger.Info.Println("Starting Maestro application...")

	logger.Info.Println("Loading environmental variables...")
	loadEnv()

	http.HandleFunc("/card/album", routeCard.AlbumHandler)
	http.HandleFunc("/card/dye", routeCard.DyeHandler)
	http.HandleFunc("/card/preview", routeCard.PreviewHandler)
	http.HandleFunc("/card/view", routeCard.ViewHandler)

	addr := os.Getenv("APP_ADDRESS")
	logger.Info.Printf("Started application on %s.\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}

func loadEnv() {
	var filename string
	flag.StringVar(&filename, "env", ".env", "switches which environmental file should be used (default: .env)")
	flag.Parse()

	if err := godotenv.Load(filename, ".env"); err != nil {
		logger.Error.Panic(err) // Unrecoverable!
	}

	if _, available := os.LookupEnv("APP_ADDRESS"); !available {
		logger.Error.Panic("environmental file is missing credentials (invalid .env file)")
	}
}
