package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lyihongl/main/snippet/app"
	//"github.com/lyihongl/main/snippet/res"
)

func SnippetRouter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println("vars", vars)
	if vars["action"] == "" {
		app.Snippet(w, r)
	} else if vars["action"] == "home" {
		app.SnippetHome(w, r)
	} else if vars["action"] == "create" {
		app.SnippetCreate(w, r)
	} else if vars["action"] == "edit" {
		app.SnippetEdit(w, r, vars["id"])
	}
}
