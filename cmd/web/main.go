package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/anucha-tk/go_booking/pkg/config"
	"github.com/anucha-tk/go_booking/pkg/handlers"
	"github.com/anucha-tk/go_booking/pkg/render"
)

var (
	app     config.AppConfig
	session *scs.SessionManager
)

func main() {
	const port = ":8080"
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

	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	render.NewTemplates(&app)
	r := routes()

	log.Println("run on port", port)
	srv := http.Server{
		Addr:              "localhost" + port,
		Handler:           r,
		ReadHeaderTimeout: time.Second * 5,
	}
	_ = srv.ListenAndServe()
}
