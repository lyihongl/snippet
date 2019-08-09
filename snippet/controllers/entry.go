package controllers

import (
	"net/http"
	"regexp"
	"text/template"

	//"github.com/lyihongl/main/snippet/data"
	"github.com/lyihongl/main/snippet/data"
	"github.com/lyihongl/main/snippet/res"
	"golang.org/x/crypto/bcrypt"
	//"github.com/lyihongl/main/snippet/data"
	//"golang.org/x/crypto/bcrypt"
	//"golang.org/x/crypto/bcrypt"
)

const (
	userNameTooLong   string = "Username cannot be over 16 characters"
	userNameEmpty            = "Username field empty"
	userAlreadyExists        = "Username already exists"
	invalidEmail             = "Invalid email address"
	passwordEmpty            = "Password field empty"
	//invalidPassword 		 = "Invalid password or username"
	invalidLogin = "Login credentials invalid"
)

//SnippetLogin serves the login page, and handles GET and POST requests
func SnippetLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles(res.VIEWS + "/snippet_login.html")
		res.CheckErr(err)
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		r.ParseForm()
		stmt, err := data.DB.Query("SELECT * FROM users WHERE username=?", r.Form.Get("username"))
		//fmt.Println(r.Form.Get("username"))
		res.CheckErr(err)
		stmt.Next()

		var uid int
		var username string
		var email string
		var password string

		stmt.Scan(&uid, &username, &email, &password)

		//fmt.Println(username, email, password)
		//fmt.Println("Username: ", r.Form["username"])
		//encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(r.Form["password"][0]), bcrypt.DefaultCost)

	}
}

//CreateErrors defines a structure to hold error states and messages during
//the creation of a user
type CreateErrors struct {
	UsernameError bool
	EmailError    bool
	PasswordError bool

	UsernameMessage []string
	EmailMessage    []string
	PasswordMessage []string

	Persist map[string]string
}

type LoginErrors struct {
	UsernameError bool
	PasswordError bool

	UsernameMessage []string
	PasswordMessage []string
}

//CreateAcc serves the create account page and handles GET and POST requests
func CreateAcc(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(res.VIEWS + "/create_acc.html")
	res.CheckErr(err)
	if r.Method == "GET" {
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		createErrors := checkCreateError(r)
		if createErrors.UsernameError || createErrors.EmailError || createErrors.PasswordError {
			t.Execute(w, createErrors)
		} else {
			stmt, err := data.DB.Prepare("INSERT INTO users (username,email,password) VALUES (?,?,?)")
			res.CheckErr(err)
			hash, err := bcrypt.GenerateFromPassword([]byte(r.Form.Get("password")), bcrypt.DefaultCost)
			res.CheckErr(err)
			stmt.Exec(r.Form.Get("username"), r.Form.Get("email"), hash)
		}
	}
}

func checkCreateError(r *http.Request) CreateErrors {
	r.ParseForm()
	var createErrors CreateErrors

	//collect values for persisting, ux thing
	createErrors.Persist = make(map[string]string)

	createErrors.Persist["username"] = r.Form["username"][0]
	createErrors.Persist["email"] = r.Form["email"][0]
	userCheck, err := data.DB.Query("SELECT * FROM users WHERE username=?", r.Form.Get("username"))
	res.CheckErr(err)
	if userCheck.Next() {
		createErrors.UsernameError = true
		createErrors.UsernameMessage = append(createErrors.UsernameMessage, userAlreadyExists)
	}

	if len(r.Form.Get("username")) > 16 {
		createErrors.UsernameError = true
		createErrors.UsernameMessage = append(createErrors.UsernameMessage, userNameTooLong)
	} else if len(r.Form.Get("username")) == 0 {
		createErrors.UsernameError = true
		createErrors.UsernameMessage = append(createErrors.UsernameMessage, userNameEmpty)
	}

	if len(r.Form.Get("password")) == 0 {
		createErrors.PasswordError = true
		createErrors.PasswordMessage = append(createErrors.PasswordMessage, passwordEmpty)
	}

	if m, _ := regexp.MatchString(`^([\w\.\_]+)@(\w+).([a-z]+)$`, r.Form.Get("email")); !m {
		createErrors.EmailError = true
		createErrors.EmailMessage = append(createErrors.EmailMessage, invalidEmail)
	}
	return createErrors
}

func checkLoginError(r *http.Request) LoginErrors{
	r.ParseForm()
	var re LoginErrors

	if len(r.Form.Get("username")) == 0 {
		re.UsernameError = true
		re.UsernameMessage = append(re.UsernameMessage, userNameEmpty)
	}

	userCheck, err := data.DB.Query("SELECT * FROM users WHERE username=?", r.Form.Get("username"))
	res.CheckErr(err)

	if !userCheck.Next() {
		re.UsernameError = true
		re.UsernameMessage = append(re.UsernameMessage, invalidLogin)
	}

	return re
}
