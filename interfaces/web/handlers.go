package web

import (
	"net/http"
)

// HomeHandler returns the home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Load the template
	tmpl := templates.Lookup("home.html")

	// Execute the template
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// AboutHandler returns the about page
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	// Load the template
	tmpl := templates.Lookup("about.html")

	// Execute the template
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ContactHandler returns the contact page
func ContactHandler(w http.ResponseWriter, r *http.Request) {
	// Load the template
	tmpl := templates.Lookup("contact.html")

	// Execute the template
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
