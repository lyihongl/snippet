package controllers

import (
	"net/http"
	"text/template"

	"github.com/lyihongl/main/snippet/res"
	"github.com/lyihongl/main/snippet/session"
)

type homeData struct {
	User string
}

func verifyUserToken() {}

//SnippetHome is the controller for the home of the application
func SnippetHome(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles(res.VIEWS + "/snippet_home.html")

		res.CheckErr(err)
		errorPage, err := template.ParseFiles(res.VIEWS + "/error.html")

		if tokenValid, user := session.ValidateToken(r); tokenValid {
			session.IssueValidationToken(w, r, user)
			test := homeData{
				User: user,
			}
			t.Execute(w, test)
		} else {
			errorPage.Execute(w, nil)
		}
	}
}
