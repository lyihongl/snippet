package app

import (
	"bytes"
	//"fmt"

	//"log"

	//"fmt"
	"net/http"
	"text/template"

	//"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/gorilla/mux"
	_data "github.com/lyihongl/main/snippet/data"
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
	var data TemplateData
	data.Init()

	errorPage, err := template.ParseFiles(res.VIEWS + "/error.gohtml")
	res.CheckErr(err)

	if r.Method == "GET" {
		t, err := template.ParseFiles(res.VIEWS + "/snippet_home.gohtml")

		res.CheckErr(err)

		data.BoolVals["logged_in"] = false

		if tokenValid, user := session.ValidateToken(r); tokenValid {

			session.IssueValidationToken(w, r, user)
			data.BoolVals["logged_in"] = true
			data.StringVals["logged_in_name"] = "Logged in as " + user
			data.StringVals["nav_bar"] = LoadTemplateAsComponenet(res.VIEWS+"/nav_bar.html", &data)

			data.StringVals["table"] = loadTableFromDB(user)

			t.Execute(w, data)

		} else {

			data.StringVals["error_msg"] = res.LOGIN_ALERT
			errorPage.Execute(w, data)

		}
	} else if r.Method == "DELETE" {
		//fmt.Println("HIT DELETE END POINT")
		id := mux.Vars(r)["id"]

		if tokenValid, _ := session.ValidateToken(r); tokenValid {
			//userid := _data.GetUserId(user)
			stmt, _ := _data.DB.Prepare("delete from snippet where id=?")	
			stmt.Exec(id)
		} else {
			data.StringVals["error_msg"] = res.LOGIN_ALERT
			errorPage.Execute(w, data)

		}


		
	}
}

func loadTableFromDB(username string) string {
	//var table string
	userid, err := _data.DB.Query("select id from users where username=?", username)
	userid.Next()
	var uid int
	userid.Scan(&uid)
	res.CheckErr(err)
	snippet, err := _data.DB.Query("SELECT id, snippet_name FROM snippet where userid=?", uid)
	res.CheckErr(err)

	var buffer bytes.Buffer

	var t_data TemplateData
	t_data.Init()
	index := 0

	for snippet.Next() {
		index++
		var uid int
		//var snippet_id int
		var name string
		//var data string

		snippet.Scan(&uid, &name)
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
	var data TemplateData
	data.Init()
	if r.Method == "GET" {
		t, err := template.ParseFiles(res.VIEWS + "/snippet_create.gohtml")

		res.CheckErr(err)
		errorPage, err := template.ParseFiles(res.VIEWS + "/error.gohtml")

		data.BoolVals["logged_in"] = false

		if tokenValid, user := session.ValidateToken(r); tokenValid {

			session.IssueValidationToken(w, r, user)
			data.BoolVals["logged_in"] = true
			data.StringVals["logged_in_name"] = "Logged in as " + user
			data.StringVals["nav_bar"] = LoadTemplateAsComponenet(res.VIEWS+"/nav_bar.html", &data)

			data.StringVals["table"] = loadTableFromDB(user)

			t.Execute(w, data)

		} else {
			data.StringVals["error_msg"] = res.LOGIN_ALERT
			errorPage.Execute(w, data)

		}
	} else if r.Method == "POST" {
		//t, err := template.ParseFiles(res.VIEWS + "/snippet_create.gohtml")

		//res.CheckErr(err)
		errorPage, err := template.ParseFiles(res.VIEWS + "/error.gohtml")
		res.CheckErr(err)

		data.BoolVals["logged_in"] = false

		if tokenValid, user := session.ValidateToken(r); tokenValid {
			r.ParseForm()
			session.IssueValidationToken(w, r, user)
			data.BoolVals["logged_in"] = true
			data.StringVals["logged_in_name"] = "Logged in as " + user
			data.StringVals["nav_bar"] = LoadTemplateAsComponenet(res.VIEWS+"/nav_bar.html", &data)

			data.StringVals["table"] = loadTableFromDB(user)

			//fmt.Println(r.Form["snippet_data"])
			//checkForSnippet, _ := _data.DB.Query("select * from snippet where name=?", r.Form["snippet_name"])
			//fmt.Println("db: ")
			//fmt.Println(_data.DB)
			//fmt.Println(user)
			userid := _data.GetUserId(user)
			//fmt.Println(userid)
			stmt, _ := _data.DB.Prepare("insert into snippet (userid, snippet_name) values (?, ?)")
			stmt.Exec(userid, r.Form.Get("snippet_name"))
			http.Redirect(w, r, "../snippet/home", 302)

			//t.Execute(w, data)
		} else {
			data.StringVals["error_msg"] = res.LOGIN_ALERT
			errorPage.Execute(w, data)

		}
	}
}

func SnippetEdit(w http.ResponseWriter, r *http.Request, id int) {
	if r.Method == "GET" {

	} else if r.Method == "POST" {

	}
}