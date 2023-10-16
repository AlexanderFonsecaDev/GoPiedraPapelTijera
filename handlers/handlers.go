package handlers

import (
	"net/http"
	"rpsweb/templates"
)

func Index(w http.ResponseWriter, r *http.Request) {
	templates.LoadTemplate(w, "index", nil)

}

func NewGame(w http.ResponseWriter, r *http.Request) {
	templates.LoadTemplate(w, "new-game", nil)
}

func Game(w http.ResponseWriter, r *http.Request) {
	templates.LoadTemplate(w, "game", nil)
}

func Play(w http.ResponseWriter, r *http.Request) {
	templates.LoadTemplate(w, "index", nil)
}

func About(w http.ResponseWriter, r *http.Request) {
	templates.LoadTemplate(w, "about", nil)
}
