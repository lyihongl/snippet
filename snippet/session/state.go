package session

import (
	"net/http"

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

	return true, claims.Username
}
