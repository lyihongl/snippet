package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	app "github.com/lyihongl/main/snippet/controllers"
	//"github.com/lyihongl/main/snippet/data"

	data "github.com/lyihongl/main/snippet/data"
	"github.com/mholt/certmagic"
	//test "github.com/lyihongl/main/test"
)

func main() {
	prod := os.Args
	fmt.Println(prod[1])
	fmt.Println(prod[1] == "prod")
	certmagic.Default.Agreed = true

	certmagic.Default.Email = "yihongliu00@gmail.com"
	data.Init()

	r := mux.NewRouter()
	r.HandleFunc("/", app.Index)

	//snippet := r.PathPrefix("/snippet").Subrouter()
	//snippet.HandleFunc("/{action}/", app.SnippetLogin)
	r.HandleFunc("/coming_soon", app.ComingSoon)
	r.HandleFunc("/login", app.GeneralLogin)
	r.HandleFunc("/create_acc", app.CreateAcc)
	r.HandleFunc("/services", app.ServiceRouter)
	r.HandleFunc("/services/{action}", app.ServiceRouter)
	//r.HandleFunc("/snippet/", app.Snippet)
	//r.HandleFunc("/snippet/{action}/", app.SnippetAction)
	//r.HandleFunc("/create_acc/", app.CreateAcc)
	r.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("snippet/views/static"))))
	r.PathPrefix("/dynamic").Handler(http.StripPrefix("/dynamic/", http.FileServer(http.Dir("snippet/javascript"))))

	if prod[1] == "prod" {
		certmagic.HTTPS([]string{"yihong.ca"}, r)
	} else {
		err := http.ListenAndServe(":9090", r) //set listen port
		if err != nil {
			log.Fatal("ListenAndServer:", err)
		}
	}

}

//func callInterface()
