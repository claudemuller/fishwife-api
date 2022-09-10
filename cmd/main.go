package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/claudemuller/fishwife/api"
	"github.com/claudemuller/fishwife/internal/pkg"
)

var (
	genQr = flag.Bool("genQr", false, "generate a QR code")
	url   = flag.String("url", "", "the url to generate into a QR code")
)

func main() {
	log.SetPrefix("__fishwife: ")
	flag.Parse()

	if *genQr {
		if *url == "" {
			log.Println("url flag is required when generating QR codes")
			flag.Usage()
			return
		}

		pkg.GenQrCode(*url)
		return
	}

	app := pkg.NewAppState()
	defer app.DB.Close()
	router := api.NewRouter(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("starting server")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
