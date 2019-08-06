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
		t, err := template.ParseFiles(res.VIEWS + "/snippet_home.html")
		res.CheckErr(err)
		t.Execute(w, nil)
	}
}

func SnippetAction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["action"] == "create_acc" {
		CreateAcc(w, r)
	}
	if vars["action"] == "login" {
		SnippetLogin(w, r)
	}
	fmt.Println("vars", vars)
}
