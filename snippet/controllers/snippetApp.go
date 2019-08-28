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

func SnippetHome(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles(res.VIEWS + "/snippet_home.html")

		res.CheckErr(err)
		errorPage, err := template.ParseFiles(res.VIEWS + "/error.html")

		if a, b := session.ValidateToken(r); a {
			test := homeData{
				User: b,
			}
			t.Execute(w, test)
		} else {
			errorPage.Execute(w, nil)
		}

		//c, err := r.Cookie("token")

		//confirm that the username has not been changed
		//if !tkn.Valid {
		//	http.Redirect(w, r, "../login/", 302)
		//} else {
		//	test := homeData{
		//		User: claims.Username,
		//	}
		//	t.Execute(w, test)
		//}

	}
}
