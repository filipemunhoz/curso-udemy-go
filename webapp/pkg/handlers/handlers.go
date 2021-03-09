package handlers

import (
	"net/http"

	"github.com/filipemunhoz/myprogram/config"
	"github.com/filipemunhoz/myprogram/models"

	"github.com/filipemunhoz/myprogram/pkg/render"
)

// Repo ...
var Repo *Repository

// Repository ...
type Repository struct {
	App *config.AppConfig
}

// NewRepo ...
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers ...
func NewHandlers(r *Repository) {
	Repo = r
}

// Home ...
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	remoteIP := r.RemoteAddr

	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})

}

// About ...
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
