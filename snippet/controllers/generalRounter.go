package controllers

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/lyihongl/main/snippet/res"
	"github.com/lyihongl/main/snippet/session"
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
	var generalData GeneralData
	generalData.LoggedIn = false
	if r.Method == "GET" {
		if a, _ := session.ValidateToken(r); a {
			generalData.LoggedIn = true
		}
		t, err := template.ParseFiles(res.VIEWS + "/services_intro.gohtml")
		res.CheckErr(err)
		t.Execute(w, generalData)
	}
}
