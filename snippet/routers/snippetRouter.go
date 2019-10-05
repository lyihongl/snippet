package controllers

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/lyihongl/main/snippet/res"
)

func Snippet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles(res.VIEWS + "/snippet_intro.gohtml")
		res.CheckErr(err)
		t.Execute(w, nil)
	}
}

func SnippetAction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println("vars", vars)
	if vars["action"] == "create_acc" {
		CreateAcc(w, r)
	} else if vars["action"] == "login" {
		GeneralLogin(w, r)
	} else if vars["action"] == "home" {
		SnippetHome(w, r)
	}
}
