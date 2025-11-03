package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/affanwhat/bookings/internal/models"
	"github.com/affanwhat/bookings/internal/config"
	"github.com/affanwhat/bookings/internal/render"
)

// Repo: the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository of the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_IP", remoteIP)

	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perfrom some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_IP")
	stringMap["remote_IP"] = remoteIP

	// send the data to the template
	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Naturale is the natural campsite page handler
func (m *Repository) Naturale(w http.ResponseWriter, r *http.Request) {
	// send the data to the template
	render.RenderTemplate(w, r, "naturale.page.tmpl", &models.TemplateData{})
}

// Cozy is the cozy suite page handler
func (m *Repository) Cozy(w http.ResponseWriter, r *http.Request) {
	// send the data to the template
	render.RenderTemplate(w, r, "cozy.page.tmpl", &models.TemplateData{})
}

// Reservation is the reservation page handler
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	// send the data to the template
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{})
}

// Contact is the contact page handler
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	// send the data to the template
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}

// Availability is the search availability page handler
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	// send the data to the template
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

// Availability is the search availability page handler
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	// post the data to the template
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("Start date is %s and end date is %s", start, end)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON is the search availability JSON page handler
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
