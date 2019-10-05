package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lyihongl/main/snippet/app"
)

//route is the general router for all services
func ServiceRouter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	print("url action: "+vars["service"]+"\n")
	if vars["service"] == "" {
		app.ServicePage(w, r)
	} else if vars["service"] == "snippet" {
		SnippetRouter(w, r)
	}
}
