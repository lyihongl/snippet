package session

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
)

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

//ConfirmUsername confirms that the username stored in browser cookies is correct,
//this function does not do any error handling
func ConfirmUsername(w http.ResponseWriter, r *http.Request) bool {
	hash, _ := r.Cookie("username_hash")
	user, _ := r.Cookie("username")
	if bcrypt.CompareHashAndPassword([]byte(hash.Value), []byte(user.Value)) != nil {
		return false
	}
	return true
}

var JwtKey = []byte("secret_key")

//ValidateToken returns true if a valid login token is stored in cookies, false otherwise
func ValidateToken(r *http.Request) (bool, string) {
	c, err := r.Cookie("token")
	if err != nil {
		return false, ""
	}
	tknStr := c.Value
	if err != nil {
		return false, ""
	}

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if tkn == nil || err != nil {
		return false, ""
	}

	if !tkn.Valid {
		return false, ""
	}
	fmt.Println("validate token username: ", claims.Username)

	return true, claims.Username
}

func IssueValidationToken(w http.ResponseWriter, r *http.Request, username string) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	fmt.Println("form username", r.Form.Get("username"))

	var _username string

	if len(r.Form.Get("username")) == 0 {
		_username = username
	} else {
		_username = r.Form.Get("username")
	}
	claims := &Claims{
		Username: _username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Path:    "/",
		Expires: expirationTime,
	})
	//user, _ := bcrypt.GenerateFromPassword([]byte(username), bcrypt.DefaultCost)

	//setting a cookie with hashed username, as well as regular username. To ensure the user is valid
	//use bcrypt to compare the 2

	//if _, err := r.Cookie("username"); err != nil {
	//http.SetCookie(w, &http.Cookie{
	//Name:    "username_hash",
	//Value:   string(user),
	//Path:    "/",
	//Expires: expirationTime,
	//})
	//}
}
