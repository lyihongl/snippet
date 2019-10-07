package app

import (
	"bytes"
	//"fmt"
	"net/http"
	"text/template"

	"github.com/lyihongl/main/snippet/data"
	"github.com/lyihongl/main/snippet/res"
	"github.com/lyihongl/main/snippet/session"
)

type homeData struct {
	User string
}

func Snippet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data TemplateData
		data.Init()

		if a, user := session.ValidateToken(r); a {
			data.BoolVals["logged_in"] = true
			data.StringVals["logged_in_name"] = "Logged in as " + user
		}
		data.StringVals["nav_bar"] = LoadTemplateAsComponenet(res.VIEWS+"/nav_bar.html", &data)

		t, err := template.ParseFiles(res.VIEWS + "/snippet_intro.gohtml")
		res.CheckErr(err)
		t.Execute(w, data)
	}
}

//SnippetHome is the controller for the home of the application
func SnippetHome(w http.ResponseWriter, r *http.Request) {
	var message res.ErrorMessage
	message.ErrorMessage = append(message.ErrorMessage, "<script>alert(You are not logged in);</script>")

	if r.Method == "GET" {
		t, err := template.ParseFiles(res.VIEWS + "/snippet_home.gohtml")

		res.CheckErr(err)
		errorPage, err := template.ParseFiles(res.VIEWS + "/error.gohtml")

		var data TemplateData
		data.Init()

		data.BoolVals["logged_in"] = false

		if tokenValid, user := session.ValidateToken(r); tokenValid {

			session.IssueValidationToken(w, r, user)
			data.BoolVals["logged_in"] = true
			data.StringVals["logged_in_name"] = "Logged in as " + user
			data.StringVals["nav_bar"] = LoadTemplateAsComponenet(res.VIEWS+"/nav_bar.html", &data)

			data.StringVals["table"] = loadTableFromDB()

			t.Execute(w, data)

		} else {

			errorPage.Execute(w, message)

		}
	}
}

func loadTableFromDB() string {
	//var table string
	snippet, err := data.DB.Query("SELECT * FROM snippet")
	res.CheckErr(err)

	var buffer bytes.Buffer

	var t_data TemplateData
	t_data.Init()
	index := 0

	for snippet.Next() {
		index++
		var uid int
		var snippet_id int
		var name string
		var data string


		snippet.Scan(&uid, &snippet_id, &name, &data)
		t_data.IntVals["snippet_id"] = uid
		t_data.IntVals["snippet_num"] = index
		t_data.StringVals["snippet_name"] = name

		//fmt.Println(name)

		t, err := template.ParseFiles(res.VIEWS + "/table.html")
		res.CheckErr(err)

		t.Execute(&buffer, t_data)
	}

	return buffer.String()
}

func SnippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

	} else if r.Method == "POST" {

	}
}
