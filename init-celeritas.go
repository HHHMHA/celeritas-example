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
	}
	err = cel.New(path)
	if err != nil {
		log.Fatal(err)
	}

	cel.InfoLog.Println("Debug mode is: ", cel.Debug)

	app := &application{
		App: cel,
	}

	return app
}
