package app

import (
	"net/http"
	"text/template"

	"github.com/lyihongl/main/snippet/res"
	"github.com/lyihongl/main/snippet/session"
)


//ServicePage loads and serves the services page
func ServicePage(w http.ResponseWriter, r *http.Request) {
	var data TemplateData
	data.Init()
	data.BoolVals["logged_in"] = false
	if r.Method == "GET" {

		if a, user := session.ValidateToken(r); a {
			data.BoolVals["logged_in"] = true
			data.StringVals["logged_in_name"] = "Logged in as "+user
		}

		data.StringVals["nav_bar"] = LoadTemplateAsComponent(res.VIEWS+"/nav_bar.html", &data)
		t, err := template.ParseFiles(res.VIEWS + "/services_intro.gohtml")
		res.CheckErr(err)
		t.Execute(w, data)
	}
}
