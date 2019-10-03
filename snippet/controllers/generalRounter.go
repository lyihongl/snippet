package controllers

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/lyihongl/main/snippet/res"
)

//route is the general router for all services
func ServiceRouter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	print(vars["action"])
	if vars["action"] == "" {
		ServicePage(w, r)
	} else if vars["action"] == "snippet" {
		Snippet(w, r)
	}
}

func ServicePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles(res.VIEWS + "/services_intro.gohtml")
		res.CheckErr(err)
		t.Execute(w, nil)
	}
}
