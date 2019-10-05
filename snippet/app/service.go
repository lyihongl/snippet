package app

import (
	"net/http"
	"text/template"

	"github.com/lyihongl/main/snippet/res"
	"github.com/lyihongl/main/snippet/session"
)

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
