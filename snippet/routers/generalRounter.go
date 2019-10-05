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
	var data TemplateData 
	data.Init()
	data.BoolVals["logged_in"] = false
	if r.Method == "GET" {
		if a, _ := session.ValidateToken(r); a {
			//generalData.LoggedIn = true
			data.BoolVals["logged_in"] = true
		}
		data.StringVals["nav_bar"] = LoadTemplateAsComponenet(res.VIEWS+"/nav_bar.html", &data)
		t, err := template.ParseFiles(res.VIEWS + "/services_intro.gohtml")
		res.CheckErr(err)
		t.Execute(w, data)
	}
}
