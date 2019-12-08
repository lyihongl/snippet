package app

//package snippet

import (
	"bytes"
	//"fmt"
	"net/http"
	"text/template"

	res "github.com/lyihongl/main/snippet/res"
	"github.com/lyihongl/main/snippet/session"
)

type GeneralData struct {
	LoggedIn bool
	NavBar   string
}


//TemplateData is a struct of bool, float, string, and int maps allowing for easy templating
type TemplateData struct{
	BoolVals map[string]bool
	FloatVals map[string]float64
	StringVals map[string]string
	IntVals map[string]int
}


//Init initializes all maps used in template data
func (t *TemplateData) Init() {
	t.BoolVals = make(map[string]bool)
	t.FloatVals = make(map[string]float64)
	t.StringVals = make(map[string]string)
	t.IntVals = make(map[string]int)
}

//Index is the main landing page of the webside, and only handles GET requests
func Index(w http.ResponseWriter, r *http.Request) {
	var data TemplateData
	data.Init()
	data.BoolVals["logged_in"] = false
	if r.Method == "GET" {
		//fmt.Println(generalData.NavBar)
		t, err := template.ParseFiles(res.VIEWS + "/index.gohtml")
		res.CheckErr(err)
		if a, user := session.ValidateToken(r); a {
			data.BoolVals["logged_in"] = true
			data.StringVals["logged_in_name"] = "Logged in as "+user
		}
		data.StringVals["nav_bar"] = LoadTemplateAsComponent(res.VIEWS+"/nav_bar.html", &data)
		t.Execute(w, data)
	}
}


//LoadTemplateAsComponent loads templates, filling in data, and returns the template as a string
func LoadTemplateAsComponent(path string, data *TemplateData) string{
	t, err := template.ParseFiles(path)
	res.CheckErr(err)
	var tpl bytes.Buffer
	t.Execute(&tpl, data)
	return tpl.String()
}


//ComingSoon handles GET requests to the coming soon pages
func ComingSoon(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles(res.VIEWS + "/coming_soon.gohtml")
		res.CheckErr(err)
		t.Execute(w, nil)
	}
}

func Projects(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles(res.VIEWS+ "/projects.html")
		res.CheckErr(err)
		t.Execute(w, nil)
	}
}