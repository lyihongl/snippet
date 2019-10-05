package app

import (
	"net/http"
	"text/template"

	"github.com/lyihongl/main/snippet/res"
	"github.com/lyihongl/main/snippet/session"
)

type homeData struct {
	User string
}

func Snippet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data TemplateData
		data.Init()

		if a, user := session.ValidateToken(r); a {
			data.BoolVals["logged_in"] = true
			data.StringVals["logged_in_name"] = "Logged in as "+user
		}
		data.StringVals["nav_bar"] = LoadTemplateAsComponenet(res.VIEWS+"/nav_bar.html", &data)

		t, err := template.ParseFiles(res.VIEWS + "/snippet_intro.gohtml")
		res.CheckErr(err)
		t.Execute(w, data)
	}
}

//SnippetHome is the controller for the home of the application
func SnippetHome(w http.ResponseWriter, r *http.Request) {
	var message res.ErrorMessage
	message.ErrorMessage = append(message.ErrorMessage, "<script>alert(You are not logged in);</script>")

	if r.Method == "GET" {
		t, err := template.ParseFiles(res.VIEWS + "/snippet_home.gohtml")

		res.CheckErr(err)
		errorPage, err := template.ParseFiles(res.VIEWS + "/error.gohtml")

		var data TemplateData
		data.Init()

		data.BoolVals["logged_in"] = false
		if tokenValid, user := session.ValidateToken(r); tokenValid {
			session.IssueValidationToken(w, r, user)
			data.BoolVals["logged_in"] = true
			data.StringVals["logged_in_name"] = "Logged in as "+user
			data.StringVals["nav_bar"] = LoadTemplateAsComponenet(res.VIEWS+"/nav_bar.html", &data)
			t.Execute(w, data)
		} else {
			errorPage.Execute(w, message)
		}
	}
}
