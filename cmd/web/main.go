package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/ferhatyegin/goBookings/internal/config"
	"github.com/ferhatyegin/goBookings/internal/handlers"
	"github.com/ferhatyegin/goBookings/internal/models"
	"github.com/ferhatyegin/goBookings/internal/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// Main is the main application function
func main() {
	//What am I going to put in the session
	gob.Register(models.Reservation{})

	// Change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache.")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	// This is where we start our web application
	fmt.Printf("starting application on port %s\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}
