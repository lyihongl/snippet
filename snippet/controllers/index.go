package controllers

//package snippet

import (
	"net/http"
	"text/template"

	res "github.com/lyihongl/main/snippet/res"
	"github.com/lyihongl/main/snippet/session"
)

type GeneralData struct{
	LoggedIn bool
}

//Index is the main landing page of the webside, and only handles GET requests
func Index(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.Method)
	var generalData GeneralData
	generalData.LoggedIn = false
	if r.Method == "GET" {
		t, err := template.ParseFiles(res.VIEWS + "/index.gohtml")
		if a, _ := session.ValidateToken(r); a {
			generalData.LoggedIn = true
		}
		res.CheckErr(err)
		t.Execute(w, generalData)
	}
}
