package controllers

import (
	"net/http"
	"text/template"

	"github.com/lyihongl/main/snippet/res"
)

func verifyUserToken() {}

func SnippetHome(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles(res.VIEWS + "/snippet_home.html")
		res.CheckErr(err)
		t.Execute(w, nil)
	}
}
