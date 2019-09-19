package main

import (
	//"fmt"
	"log"
	"net/http"
	"crypto/tls"
	"golang.org/x/crypto/acme/autocert"

	"github.com/gorilla/mux"
	app "github.com/lyihongl/main/snippet/controllers"
	data "github.com/lyihongl/main/snippet/data"
	//test "github.com/lyihongl/main/test"
)

func main() {

	certManager := autocert.Manager{
		Prompt:		autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("yihong.ca"),
		Cache:		autocert.DirCache("certs"),
	}

	//go globalSessions.GC
	data.Init()
	data.GetConfig("./snippet/data/env.txt")

	r := mux.NewRouter()
	r.HandleFunc("/", app.Index)

	//snippet := r.PathPrefix("/snippet").Subrouter()
	//snippet.HandleFunc("/{action}/", app.SnippetLogin)
	r.HandleFunc("/login/", app.GeneralLogin)
	r.HandleFunc("/create_acc/", app.CreateAcc)

	r.HandleFunc("/snippet/", app.Snippet)
	r.HandleFunc("/snippet/{action}/", app.SnippetAction)
	//r.HandleFunc("/create_acc/", app.CreateAcc)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("snippet/views/static"))))
	r.PathPrefix("/dynamic/").Handler(http.StripPrefix("/dynamic/", http.FileServer(http.Dir("snippet/javascript"))))

	server := &http.Server{
		Addr: ":https",
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
		Handler: r,
	}

	go http.ListenAndServe(":http", certManager.HTTPHandler(nil))

	log.Fatal(server.ListenAndServeTLS("", ""))

	err := http.ListenAndServe(":443", r) //set listen port
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}

//func callInterface()
