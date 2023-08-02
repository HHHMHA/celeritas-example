package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"myapp/data"
	"net/http"
)

func (a *application) routes() *chi.Mux {
	// middleware must come before any routes

	// add routes here
	a.App.Routes.Get("/", a.Handlers.Home)
	a.App.Routes.Get("/go-page", a.Handlers.GoPage)
	a.App.Routes.Get("/jet-page", a.Handlers.JetPage)
	a.App.Routes.Get("/sessions", a.Handlers.SessionTest)
	//a.App.Routes.Get("/sessions", a.Handlers.Home)
	a.App.Routes.Get("/create-user", func(w http.ResponseWriter, r *http.Request) {
		u := data.User{
			FirstName: "Trevor",
			LastName:  "Sawler",
			Email:     "me@here.com",
			Active:    1,
			Password:  "password",
		}

		id, err := a.Models.Users.Insert(u)
		if err != nil {
			a.App.ErrorLog.Println(err)
			return
		}

		fmt.Fprintf(w, "%d: %s", id, u.FirstName)
	})

	a.App.Routes.Get("/test-database", func(w http.ResponseWriter, r *http.Request) {
		query := "SELECT id, first_name FROM users WHERE id = 1"
		row := a.App.DB.Pool.QueryRowContext(r.Context(), query)

		var id int
		var name string
		err := row.Scan(&id, &name)

		if err != nil {
			a.App.ErrorLog.Println(err)
			return
		}

		fmt.Fprintf(w, "%d %s", id, name)
	})

	// static routes
	fileServer := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return a.App.Routes
}
