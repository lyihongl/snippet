package controllers

//package snippet

import (
	"net/http"
	"text/template"

	res "github.com/lyihongl/main/snippet/res"
)

//Index is the main landing page of the webside, and only handles GET requests
func Index(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles(res.VIEWS + "/index.html")
		res.CheckErr(err)
		t.Execute(w, nil)
	}
}
