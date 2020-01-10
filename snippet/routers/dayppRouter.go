package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lyihongl/main/snippet/app"
)

func DayPPRouter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["action"] == "" {
		app.DayPP(w, r)
	}
}
