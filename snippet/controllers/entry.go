package controllers

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

const (
	userNameTooLong   = "Username cannot be over 16 characters"
	userNameEmpty     = "Username field empty"
	userAlreadyExists = "Username already exists"
	invalidEmail      = "Invalid email address"
	passwordEmpty     = "Password field empty"
	invalidLogin      = "Login credentials invalid"
	passwordMissmatch = "Passwords do not match"
)

type Test struct{
	Test bool
}

//SnippetLogin serves the login page, and handles GET and POST requests
func GeneralLogin(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(res.VIEWS + "/general_login.html")
	if r.Method == "GET" {
		if a, _ := session.ValidateToken(r); a {
			http.Redirect(w, r, "../", 200)
		}
		res.CheckErr(err)
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		r.ParseForm()
		loginErrors := checkLoginError(r)
		if loginErrors.UsernameError || loginErrors.PasswordError {
			t.Execute(w, loginErrors)
		} else {

			session.IssueValidationToken(w, r, r.Form.Get("username"))
			http.Redirect(w, r, "../", 200)
			//http.SetCookie(w, &http.Cookie{
			//	Name:		"username",
			//	Value:		r.Form.Get("username"),
			//	Expires:	expirationTime,
			//})
		}

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

//LoginErrors is a struct that holds login error state
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
			http.Redirect(w, r, "..", http.StatusFound)
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

	//var creds session.Credentials
	//err := json.NewDecoder(r.Body).Decode(&creds)
	//res.CheckErr(err)

	if len(r.Form.Get("username")) == 0 {
		re.UsernameError = true
		re.UsernameMessage = append(re.UsernameMessage, userNameEmpty)
	}

	userCheck, err := data.DB.Query("SELECT * FROM users WHERE username=?", r.Form.Get("username"))
	res.CheckErr(err)

	if !userCheck.Next() {
		fail = true
	}

	var uid int
	var username string
	var email string
	var password string

	userCheck.Scan(&uid, &username, &email, &password)

	//fmt.Println([]byte(password))
	//test := bcrypt.CompareHashAndPassword([]byte(password), []byte("Rubixcube123"))
	//fmt.Println(test)

	if bcrypt.CompareHashAndPassword([]byte(password), []byte(r.Form.Get("password"))) != nil {
		fail = true
	}
	if fail {
		re.UsernameError = true
		re.UsernameMessage = append(re.UsernameMessage, invalidLogin)
	}
	return re
}
