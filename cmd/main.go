package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/claudemuller/fishwife/api"
	"github.com/claudemuller/fishwife/internal/pkg"
)

var mode = flag.Bool("serve", true, "run in server mode")

func main() {
	log.SetPrefix("__fishwife: ")

	app := pkg.NewAppState()
	defer app.DB.Close()
	router := api.NewRouter(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("starting server")
	log.Fatal(http.ListenAndServe("localhost:"+port, router))
}
