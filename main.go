package main

import (
	//"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	routers "github.com/lyihongl/main/snippet/routers"
	app "github.com/lyihongl/main/snippet/app"
	//"github.com/lyihongl/main/snippet/data"

	//data "github.com/lyihongl/main/snippet/data"
	"github.com/mholt/certmagic"
	//test "github.com/lyihongl/main/test"
)

func main() {
	prod := os.Args
	//fmt.Println(prod[1])
	//fmt.Println(prod[1] == "prod")

	certmagic.Default.Agreed = true
	certmagic.Default.Email = "yihongliu00@gmail.com"

	//init database
	//data.Init()

	r := mux.NewRouter()
	r.HandleFunc("/", app.Index)

	//General routes
	r.HandleFunc("/coming_soon", app.ComingSoon)
	r.HandleFunc("/login", app.GeneralLogin)
	r.HandleFunc("/create_acc", app.CreateAcc)
	r.HandleFunc("/projects", app.Projects)

	//service routes
	r.HandleFunc("/services", routers.ServiceRouter)
	r.HandleFunc("/services/{service}", routers.ServiceRouter)
	r.HandleFunc("/services/{service}/{action}", routers.ServiceRouter)
	r.HandleFunc("/services/{service}/{action}/{id}", routers.ServiceRouter)


	r.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("snippet/static"))))
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