package controllers

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

type TemplateData struct{
	BoolVals map[string]bool
	FloatVals map[string]float64
	StringVals map[string]string
}

func InitTemplateData(t *TemplateData) {
	t.BoolVals = make(map[string]bool)
	t.FloatVals = make(map[string]float64)
	t.StringVals = make(map[string]string)
}

//Index is the main landing page of the webside, and only handles GET requests
func Index(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.Method)
	//var generalData GeneralData
	//generalData.LoggedIn = false
	var data TemplateData
	data.BoolVals["logged_in"] = false
	if r.Method == "GET" {
		//fmt.Println(generalData.NavBar)
		t, err := template.ParseFiles(res.VIEWS + "/index.gohtml")
		res.CheckErr(err)
		if a, _ := session.ValidateToken(r); a {
			data.BoolVals["logged_in"] = true
		}
		data.StringVals["nav_bar"] = LoadTemplateAsComponenet(res.VIEWS+"/nav_bar.html", &data)
		t.Execute(w, data)
	}
}

func LoadTemplateAsComponenet(path string, data *TemplateData) string{
	t, err := template.ParseFiles(path)
	res.CheckErr(err)
	var tpl bytes.Buffer
	t.Execute(&tpl, data)
	return tpl.String()
}

func ComingSoon(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles(res.VIEWS + "/coming_soon.gohtml")
		res.CheckErr(err)
		t.Execute(w, nil)
	}
}
