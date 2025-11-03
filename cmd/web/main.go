package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/affanwhat/bookings/internal/config"
	"github.com/affanwhat/bookings/internal/handlers"
	"github.com/affanwhat/bookings/internal/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8081"
var app config.AppConfig
var session *scs.SessionManager

func main() {
	

	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	render.NewTemplates(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	/*
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	*/

	fmt.Println(fmt.Sprintf("Starting application of port %s \n", portNumber))
	//_ = http.ListenAndServe(portNumber, nil) // If there's an error, ignore it.

	srv := &http.Server {
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
