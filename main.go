package main

import (
	"github.com/HHHMHA/celeritas"
	"myapp/data"
	"myapp/handlers"
)

type application struct {
	App      *celeritas.Celeritas
	Handlers *handlers.Handlers
	Models   data.Models
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
