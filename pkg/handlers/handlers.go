package handlers

import (
	"net/http"
	"github.com/georgeikani/hotel-booking/pkg/render"
)


// This is the home page
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.html")
}


// This is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.html")
}


