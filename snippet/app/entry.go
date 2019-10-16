package app

import (
	//"fmt"
	"net/http"
	"regexp"
	"text/template"

	"github.com/lyihongl/main/snippet/data"
	"github.com/lyihongl/main/snippet/res"
	"github.com/lyihongl/main/snippet/session"
	"golang.org/x/crypto/bcrypt"
)

//Error messages
const (
	userNameTooLong   	= "Username cannot be over 16 characters"
	userNameEmpty     	= "Username field empty"
	userAlreadyExists 	= "Username already exists"
	invalidEmail      	= "Invalid email address"
	emailAlreadyExists 	= "Account with this email already exists"
	passwordEmpty     	= "Password field empty"
	invalidLogin      	= "Login credentials invalid"
	passwordMissmatch 	= "Passwords do not match"
)

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

//LoginErrors is a struct that holds login error state
type LoginErrors struct {
	UsernameError bool
	PasswordError bool

	UsernameMessage []string
	PasswordMessage []string
}

//GeneralLogin serves the login page, and handles GET and POST requests
func GeneralLogin(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(res.VIEWS + "/general_login.gohtml")
	if r.Method == "GET" {
		if a, _ := session.ValidateToken(r); a {
			http.Redirect(w, r, "../", 302)
		} else {
			res.CheckErr(err)
			t.Execute(w, nil)
		}
	} else if r.Method == "POST" {
		r.ParseForm()
		loginErrors := checkLoginError(r)
		if loginErrors.UsernameError || loginErrors.PasswordError {
			t.Execute(w, loginErrors)
		} else {
			session.IssueValidationToken(w, r, r.Form.Get("username"))
			http.Redirect(w, r, "/", 302)
		}
	}
}


//CreateAcc serves the create account page and handles GET and POST requests
func CreateAcc(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(res.VIEWS + "/create_acc.gohtml")
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
			http.Redirect(w, r, "..", http.StatusFound)
		}
	}
}

type userInterface struct{
	uid int
	username string
	email string
	password string
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

	emailCheck, err := data.DB.Query("select * from users where email=?", r.Form.Get("email"))
	res.CheckErr(err)

	if emailCheck.Next() {
		createErrors.EmailError = true
		createErrors.EmailMessage = append(createErrors.EmailMessage, emailAlreadyExists)
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

	if r.Form.Get("password") != r.Form.Get("confirm_password") {
		createErrors.PasswordError = true
		createErrors.PasswordMessage = append(createErrors.PasswordMessage, passwordMissmatch)
	}
	return createErrors
}

func checkLoginError(r *http.Request) LoginErrors {
	r.ParseForm()
	var re LoginErrors

	fail := false

	if len(r.Form.Get("username")) == 0 {
		re.UsernameError = true
		re.UsernameMessage = append(re.UsernameMessage, userNameEmpty)
	}

	userCheck, err := data.DB.Query("SELECT password FROM users WHERE username=?", r.Form.Get("username"))
	res.CheckErr(err)

	if !userCheck.Next() {
		fail = true
	}

	//var uid int
	//var username string
	//var email string
	var password string

	userCheck.Scan(&password)

	if bcrypt.CompareHashAndPassword([]byte(password), []byte(r.Form.Get("password"))) != nil {
		fail = true
	}
	if fail {
		re.UsernameError = true
		re.UsernameMessage = append(re.UsernameMessage, invalidLogin)
	}
	return re
}
