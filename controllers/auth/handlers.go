package auth

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/database"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/helpers"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/helpers/password"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/models"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/repository/users"

	"github.com/dgrijalva/jwt-go"
)

//JwtKey is the JSONWebToken key used for auth
var JwtKey = []byte("my_secret_key")

//Credentials are the credentials used for auth
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

//GoQRSClaims is a claims model
type GoQRSClaims struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

//Signin is used for signing in
func Signin(w http.ResponseWriter, r *http.Request) {

	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db := database.DbConn
	repository := users.Repository{Conn: db}

	user, err := repository.GetUser(creds.Username)
	if err != nil {
		log.Printf("could not get auth: %v", err)
		return
	}
	if user == nil {
		log.Print("no auth found")
		helpers.WriteErrorJSON(w, http.StatusBadRequest, "no auth to connect")
		return
	}
	isMatching, err := password.ComparePasswordAndHash(creds.Password, user.Password)
	if err != nil {
		log.Printf("could not compare password: %v", err)
		helpers.WriteErrorJSON(w, http.StatusInternalServerError, "could not compare password")
		return
	}
	if !isMatching {
		log.Print("password not matching")
		helpers.WriteErrorJSON(w, http.StatusUnauthorized, "invalid credentials")
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &GoQRSClaims{
		ID:       user.ID,
		Username: creds.Username,
		Role:     user.Role,
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
		Expires: expirationTime,
	})

	helpers.WriteJSON(w, http.StatusOK, "auth logged in")
}

//SignUp is used for signing up
func SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Print(err)
		helpers.WriteErrorJSON(w, http.StatusInternalServerError, "could not decode request body")
		return
	}

	db := database.DbConn
	repository := users.Repository{Conn: db}

	userFromDB, err := repository.GetUser(user.Username)
	if err != nil {
		log.Print(err)
		helpers.WriteErrorJSON(w, http.StatusInternalServerError, "could not get auth from db")
		return
	}

	if userFromDB != nil {
		log.Print(err)
		helpers.WriteErrorJSON(w, http.StatusBadRequest, "auth already exists")
		return
	}

	hash, err := password.GenerateFromPassword(user.Password)
	if err != nil {
		log.Print(err)
		helpers.WriteErrorJSON(w, http.StatusInternalServerError, "could not safely save auth")
		return
	}
	user.Password = hash

	err = repository.SaveUser(&user)
	if err != nil {
		log.Print(err)
		helpers.WriteErrorJSON(w, http.StatusInternalServerError, "could not save auth in db")
		return
	}

	helpers.WriteJSON(w, http.StatusOK, "auth registered")

}

//Refresh is used for refreshing the token
func Refresh(w http.ResponseWriter, r *http.Request) {
	// (BEGIN) The code uptil this point is the same as the first part of the `Welcome` route
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tknStr := c.Value
	claims := &GoQRSClaims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(2 * time.Hour)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
