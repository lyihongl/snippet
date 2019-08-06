package main

import (
	//"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	app "github.com/lyihongl/main/snippet/controllers"
	db "github.com/lyihongl/main/snippet/data"
	//test "github.com/lyihongl/main/test"
)

func main() {

	//go globalSessions.GC
	db.Init()

	r := mux.NewRouter()
	r.HandleFunc("/", app.Index)

	//snippet := r.PathPrefix("/snippet").Subrouter()
	//snippet.HandleFunc("/{action}/", app.SnippetLogin)

	r.HandleFunc("/snippet/", app.Snippet)
	r.HandleFunc("/snippet/{action}/", app.SnippetAction)
	//r.HandleFunc("/create_acc/", app.CreateAcc)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("snippet/views"))))
	//r.PathPrefix("/dynamic/").Handler(http.StripPrefix("/dynamic/", http.FileServer(http.Dir("snippet/javascript"))))

	err := http.ListenAndServe(":9090", r) //set listen port
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}

//func callInterface()
