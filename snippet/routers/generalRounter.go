package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lyihongl/main/snippet/app"
)

//ServiceRouter is the general router for all services
func ServiceRouter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars)
	fmt.Println(r.Method)

	//var snippetCrumbs BreadCrumbList

	//var levels = map[string]int{
	//	"home":ROOT,
	//	"edit":LEVEL_1,
	//	"create":LEVEL_1,
	//}
	//r.Path

	//snippetCrumbs.DefineLevels(levels)
	//breadCrumb := BreadCrumb{
	//	Title : "test",
	//	Path : r.URL.Path,
	//	Level : levels[vars["action"]],
	//}
	//snippetCrumbs.Update(vars["action"])
	//snippetCrumbs.Push(&breadCrumb)

	print("url action: " + vars["service"] + "\n")
	//fmt.Println("bread crumbs: ")
	//fmt.Println(snippetCrumbs.Crumbs)
	if vars["service"] == "" {
		app.ServicePage(w, r)
	} else if vars["service"] == "snippet" {
		SnippetRouter(w, r)
	}
}
