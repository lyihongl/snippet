package app

import (
	"bytes"
	"fmt"
	"net/http"
	"regexp"
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

//Snippet serves the page that introduces what snippet is
func Snippet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data TemplateData
		data.Init()

		if a, user := session.ValidateToken(r); a {
			data.BoolVals["logged_in"] = true
			data.StringVals["logged_in_name"] = "Logged in as " + user
		}
		data.StringVals["nav_bar"] = LoadTemplateAsComponent(res.VIEWS+"/nav_bar.html", &data)

		t, err := template.ParseFiles(res.VIEWS + "/snippet_intro.gohtml")
		res.CheckErr(err)
		t.Execute(w, data)
	}
}

//SnippetHome is the controller for the home of the application
func SnippetHome(w http.ResponseWriter, r *http.Request) {

	var data TemplateData
	data.Init()

	data.BoolVals["logged_in"] = false
	data.StringVals["nav_bar"] = LoadTemplateAsComponent(res.VIEWS+"/nav_bar.html", &data)

	if tokenValid, user := session.ValidateToken(r); tokenValid {
		data.BoolVals["logged_in"] = true
		data.StringVals["nav_bar"] = LoadTemplateAsComponent(res.VIEWS+"/nav_bar.html", &data)

		t, err := template.ParseFiles(res.VIEWS + "/snippet_home.gohtml")
		res.CheckErr(err)

		if r.Method == "GET" {
			session.IssueValidationToken(w, r, user)
			data.BoolVals["logged_in"] = true
			data.StringVals["logged_in_name"] = "Logged in as " + user

			data.StringVals["table"] = loadTableFromDB(user)

			t.Execute(w, data)
		} else if r.Method == "DELETE" {
			id := mux.Vars(r)["id"]
			stmt, _ := _data.DB.Prepare("delete from snippet where id=?")
			stmt.Exec(id)
		}
	} else {
		//TODO: add navbar and erorr message to error page
		errorPage, err := template.ParseFiles(res.VIEWS + "/error.gohtml")
		res.CheckErr(err)
		data.StringVals["error_msg"] = res.LOGIN_ALERT
		errorPage.Execute(w, data)
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

	var tData TemplateData
	tData.Init()
	index := 0

	for snippet.Next() {
		index++
		var uid int
		//var snippet_id int
		var name string
		//var data string

		snippet.Scan(&uid, &name)
		tData.IntVals["snippet_id"] = uid
		tData.IntVals["snippet_num"] = index
		tData.StringVals["snippet_name"] = name

		//fmt.Println(name)

		t, err := template.ParseFiles(res.VIEWS + "/table.html")
		res.CheckErr(err)

		t.Execute(&buffer, tData)
	}

	return buffer.String()
}

func loadExportList(username string) string {
	userid, err := _data.DB.Query("select id from users where username=?", username)
	userid.Next()
	var uid int
	userid.Scan(&uid)
	res.CheckErr(err)
	snippet, err := _data.DB.Query("SELECT id, snippet_name FROM snippet where userid=?", uid)
	res.CheckErr(err)

	var buffer bytes.Buffer

	var tData TemplateData
	tData.Init()
	index := 0
	for snippet.Next() {
		index++
		var uid int
		//var snippet_id int
		var name string
		//var data string

		snippet.Scan(&uid, &name)
		tData.IntVals["snippet_id"] = uid
		tData.IntVals["snippet_num"] = index
		tData.StringVals["snippet_name"] = name

		//fmt.Println(name)

		t, err := template.ParseFiles(res.VIEWS + "/export_table.gohtml")
		res.CheckErr(err)

		t.Execute(&buffer, tData)
	}
	return buffer.String()
}

//SnippetCreate handles requests for the snippet create page
func SnippetCreate(w http.ResponseWriter, r *http.Request) {
	var data TemplateData
	data.Init()

	data.BoolVals["logged_in"] = false

	if tokenValid, user := session.ValidateToken(r); tokenValid {

		session.IssueValidationToken(w, r, user)

		if r.Method == "GET" {
			//t, err := template.ParseFiles(res.VIEWS + "/snippet_create.gohtml")
			var t *template.Template
			t, data := LoadStdPage(r, "/snippet_create.gohtml", user)

			data.StringVals["table"] = loadTableFromDB(user)

			t.Execute(w, data)
		} else if r.Method == "POST" {

			r.ParseForm()
			session.IssueValidationToken(w, r, user)

			data.BoolVals["logged_in"] = true
			data.StringVals["logged_in_name"] = "Logged in as " + user
			data.StringVals["nav_bar"] = LoadTemplateAsComponent(res.VIEWS+"/nav_bar.html", &data)

			data.StringVals["table"] = loadTableFromDB(user)
			userid := _data.GetUserId(user)
			stmt, _ := _data.DB.Prepare("insert into snippet (userid, snippet_name) values (?, ?)")
			stmt.Exec(userid, r.Form.Get("snippet_name"))

			http.Redirect(w, r, "../snippet/home", 302)
		}
	} else {
		errorPage, err := template.ParseFiles(res.VIEWS + "/error.gohtml")
		data.StringVals["error_msg"] = res.LOGIN_ALERT
		res.CheckErr(err)
		errorPage.Execute(w, data)
	}
}

//SnippetEdit handles editing of snippets
func SnippetEdit(w http.ResponseWriter, r *http.Request, id string) {
	var data TemplateData
	data.Init()

	if tokenValid, user := session.ValidateToken(r); tokenValid {
		if r.Method == "GET" {
			t, data := LoadStdPage(r, "/snippet_edit.gohtml", user)

			nameQuery, _ := _data.DB.Query("select snippet_name, data from snippet where id=?", id)
			nameQuery.Next()

			var snippetName string
			var snippetData string
			nameQuery.Scan(&snippetName, &snippetData)

			data.StringVals["snippet_name"] = snippetName
			data.StringVals["snippet_data"] = snippetData

			stmt, _ := _data.DB.Prepare("update snippet set snippet_name=?, data=? where id=?")

			stmt.Exec(data.StringVals["snippet_name"], data.StringVals["snippet_data"], id)

			re := regexp.MustCompile(`\r?\n`)

			snippetPreview := LoadTemplateAsComponent(res.VIEWS+"/preview.html", data)

			data.StringVals["snippet_preview"] = re.ReplaceAllString(snippetPreview, "")

			previewScript := LoadTemplateAsComponent(res.VIEWS+"/preview_script.html", data)
			data.StringVals["preview_script"] = re.ReplaceAllString(previewScript, "")

			t.Execute(w, data)
		} else if r.Method == "POST" {
			r.ParseForm()
			t, data := LoadStdPage(r, "/snippet_edit.gohtml", user)
			data.StringVals["snippet_name"] = r.Form.Get("snippet_name")
			data.StringVals["snippet_data"] = r.Form.Get("snippet_data")

			stmt, _ := _data.DB.Prepare("update snippet set snippet_name=?, data=? where id=?")
			stmt.Exec(data.StringVals["snippet_name"], data.StringVals["snippet_data"], id)

			re := regexp.MustCompile(`\r?\n`)

			snippetPreview := LoadTemplateAsComponent(res.VIEWS+"/preview.html", data)

			data.StringVals["snippet_preview"] = re.ReplaceAllString(snippetPreview, "")

			previewScript := LoadTemplateAsComponent(res.VIEWS+"/preview_script.html", data)
			data.StringVals["preview_script"] = re.ReplaceAllString(previewScript, "")

			t.Execute(w, data)
		}
	}
}

//SnippetExport is the controller for the export page
func SnippetExport(w http.ResponseWriter, r *http.Request) {
	if tokenValid, user := session.ValidateToken(r); tokenValid {
		t, data := LoadStdPage(r, "/snippet_export.gohtml", user)
		fmt.Println(loadExportList(user))
		data.StringVals["export_table"] = loadExportList(user)
		t.Execute(w, data)
	} else {

	}
}

//LoadStdPage loads a page with navbar
func LoadStdPage(r *http.Request, templatePath string, user string) (*template.Template, *TemplateData) {
	var data TemplateData
	data.Init()
	t, err := template.ParseFiles(res.VIEWS + templatePath)
	res.CheckErr(err)

	fmt.Println("edit func")
	data.BoolVals["logged_in"] = true
	data.StringVals["logged_in_name"] = "Logged in as " + user
	data.StringVals["nav_bar"] = LoadTemplateAsComponent(res.VIEWS+"/nav_bar.html", &data)

	return t, &data
}
