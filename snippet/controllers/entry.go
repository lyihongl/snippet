package controllers

import (
	"fmt"
	"net/http"
	"regexp"
	"text/template"

	"github.com/lyihongl/main/snippet/res"
	//"github.com/lyihongl/main/snippet/data"
	//"golang.org/x/crypto/bcrypt"
	//"golang.org/x/crypto/bcrypt"
)

const (
	userNameTooLong string = "Username over 16 characters"
	userNameEmpty          = "Username cannot be empty"
	invalidEmail           = "Invalid email address"
	passwordEmpty 		   = "Password cannot be empty"
)

//SnippetLogin serves the login page, and handles GET and POST requests
func SnippetLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles(res.VIEWS + "/snippet_login.html")
		res.CheckErr(err)
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		r.ParseForm()
		fmt.Println("Username: ", r.Form["username"])
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

//CreateAcc serves the create account page and handles GET and POST requests
func CreateAcc(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(res.VIEWS + "/create_acc.html")
	res.CheckErr(err)
	if r.Method == "GET" {
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		createErrors := checkCreateError(r)
		if createErrors.UsernameError || createErrors.EmailError || createErrors.PasswordError {
			//fmt.Println(createErrors.UsernameError)
			//fmt.Println(createErrors.Persist)
			t.Execute(w, createErrors)
		} else {
			stmt, err := data.DB.Prepare("INSERT INTO users")
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
