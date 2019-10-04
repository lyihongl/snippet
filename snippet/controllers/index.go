package controllers

//package snippet

import (
	"bytes"
	"fmt"
	"net/http"
	"text/template"

	res "github.com/lyihongl/main/snippet/res"
	"github.com/lyihongl/main/snippet/session"
)

type GeneralData struct {
	LoggedIn bool
	NavBar   string
}

//Index is the main landing page of the webside, and only handles GET requests
func Index(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.Method)
	var generalData GeneralData
	generalData.LoggedIn = false
	if r.Method == "GET" {
		//fmt.Println(generalData.NavBar)
		t, err := template.ParseFiles(res.VIEWS + "/index.gohtml")
		if a, _ := session.ValidateToken(r); a {
			generalData.LoggedIn = true
		}
		data, _ := template.ParseFiles(res.VIEWS + "/nav_bar.html")
		var tpl bytes.Buffer
		data.Execute(&tpl, generalData)
		result := tpl.String()
		fmt.Println(result)
		generalData.NavBar = result
		res.CheckErr(err)
		t.Execute(w, generalData)
	}
}

func ComingSoon(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles(res.VIEWS + "/coming_soon.gohtml")
		res.CheckErr(err)
		t.Execute(w, nil)
	}
}
