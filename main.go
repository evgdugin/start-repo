package main

import (
	"flag"
	"log"
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
