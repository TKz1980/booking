package handlers

import (
	"net/http"

	"github.com/TKz1980/booking/pkg/config"
	"github.com/TKz1980/booking/pkg/models"
	"github.com/TKz1980/booking/pkg/render"
)



var Repo * Repository

type Repository struct {
	App *config.AppConfig

}

func NewRepo(a *config.AppConfig) * Repository {
	return &Repository {
		App: a,
	}
}

//NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the about page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote=IP", remoteIP)
	render.RenderTemplate(w,"home.page.html", &models.TemplateData{})
}


// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hollo, again!!!!"

	remoteIP := m.App.Session.GetString(r.Context(), "remote=IP")
	stringMap["remote=IP"] = remoteIP


	render.RenderTemplate(w,"about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
