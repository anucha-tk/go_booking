package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/anucha-tk/go_booking/pkg/config"
	"github.com/anucha-tk/go_booking/pkg/models"
	"github.com/anucha-tk/go_booking/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{}
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send data to the template
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Time(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "time.page.html", &models.TemplateData{})
}

func (m *Repository) GetTime(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format("15:04:05")
	_, err := w.Write([]byte(currentTime))
	if err != nil {
		log.Fatal(err)
	}
}
