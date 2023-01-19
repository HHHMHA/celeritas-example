package main

import (
	"log"
	"os"

	"github.com/HHHMHA/celeritas"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	cel := &celeritas.Celeritas{
		AppName: "myapp",
		Debug:   true,
	}
	err = cel.New(path)
	if err != nil {
		log.Fatal(err)
	}

	app := &application{
		App: cel,
	}

	return app
}
