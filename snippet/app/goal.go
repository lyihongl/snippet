package app

import (
	"bytes"
	"fmt"
	"net/http"
	"text/template"

	//"github.com/lyihongl/main/snippet/data"
	_data "github.com/lyihongl/main/snippet/data"
	"github.com/lyihongl/main/snippet/res"
	"github.com/lyihongl/main/snippet/session"
)

func loadInitiativeTable(username string) string {
	//var table string
	userid, err := _data.DB.Query("select id from users where username=?", username)
	userid.Next()
	var uid int
	userid.Scan(&uid)
	res.CheckErr(err)
	//snippet, err := _data.DB.Query("SELECT id, snippet_name FROM snippet where userid=?", uid)
	initiative, err := _data.DB.Query("SELECT id, title from initiative where userid=?", uid)
	res.CheckErr(err)

	var buffer bytes.Buffer

	var tData TemplateData
	tData.Init()
	index := 0

	for initiative.Next() {
		index++
		var id int
		var title string
		initiative.Scan(&id, &title)
		tData.IntVals["initiative_id"] = id
		tData.StringVals["title"] = title

		t, _ := template.ParseFiles(res.VIEWS + "/initiative_table.html")
		t.Execute(&buffer, tData)
	}

	return buffer.String()
}

func DayPP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles(res.VIEWS + "/DayPP.gohtml")
		res.CheckErr(err)
		// load table from db
		var data TemplateData
		data.Init()
		_, user := session.ValidateToken(r)
		data.StringVals["table"] = loadInitiativeTable(user)
		t.Execute(w, data)
	}
}

func DayPPCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles(res.VIEWS + "/daypp_create.html")
		res.CheckErr(err)

		t.Execute(w, nil)
	} else if r.Method == "POST" {
		// Insert new initiative to db
		r.ParseForm()
		//fmt.Println(r.Form.Get("title"))
		_, user := session.ValidateToken(r)
		stmt, _ := _data.DB.Prepare("insert into initiative (title, userid) values (?, ?)")
		stmt.Exec(r.Form.Get("title"), _data.GetUserId(user))

		http.Redirect(w, r, "../DayPP", 302)
	}
}

func DayPPView(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := r.URL.Query().Get("id")
		//fmt.Println(id)
		var data TemplateData
		data.Init()
		title, _ := _data.DB.Query("select title from initiative where id=?", id)
		var _title string
		if title.Next() {
			title.Scan(&_title)
		}
		fmt.Println(_title)

		data.StringVals["title"] = _title
	} else if r.Method == "POST" {

	}
}
