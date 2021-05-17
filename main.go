package main

import (
	"flag"
	"log"

	"github.com/evgdugin/start-repo/app"
)

func main() {
	// TODO:
	configFileName := flag.String("c", "config.yml", "Config")
	flag.Parse()

	if app, err := app.New(*configFileName); err != nil {
		log.Fatal(err)
	} else if err = app.Run(); err != nil {
		log.Fatal(err)
	}

}
