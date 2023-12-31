package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/TKz1980/booking/internal/config"
	"github.com/TKz1980/booking/internal/handlers"
	"github.com/TKz1980/booking/internal/models"
	"github.com/TKz1980/booking/internal/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	err := run()
	if err != nil {
		log.Fatal((err))
	}

	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	// _ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)


}

func run() error {
//what am i going to put in session
gob.Register(models.Reservation{})
		


//chang this to true when in production

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
	return err
}

app.TemplateCache = tc
app.UseCache = false

repo := handlers.NewRepo(&app)
handlers.NewHandlers(repo)

render.NewTemplates(&app)

	return nil
}