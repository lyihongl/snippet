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

		if a, _ := session.ValidateToken(r); a {
			data.BoolVals["logged_in"] = true
		}
		data.StringVals["nav_bar"] = LoadTemplateAsComponenet(res.VIEWS+"/nav_bar.html", &data)

		t, err := template.ParseFiles(res.VIEWS + "/snippet_intro.gohtml")
		res.CheckErr(err)
		t.Execute(w, data)
	}
}

//SnippetHome is the controller for the home of the application
func SnippetHome(w http.ResponseWriter, r *http.Request) {
	//message := res.ErrorMessage
	var message res.ErrorMessage
	//message.ErrorMessage = append(message.ErrorMessage, "Login token invalid, make sure cookies are enabled and try logging in again")
	message.ErrorMessage = append(message.ErrorMessage, "<script>alert(You are not logged in);</script>")
	if r.Method == "GET" {
		t, err := template.ParseFiles(res.VIEWS + "/snippet_home.gohtml")

		res.CheckErr(err)
		errorPage, err := template.ParseFiles(res.VIEWS + "/error.gohtml")

		if tokenValid, user := session.ValidateToken(r); tokenValid {
			session.IssueValidationToken(w, r, user)
			//fmt.Println(user)
			info := homeData{
				User: user,
			}
			t.Execute(w, info)
		} else {
			errorPage.Execute(w, message)
		}
	}
}
