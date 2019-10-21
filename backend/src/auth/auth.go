package auth

import "net/http"

func IsLoggedIn(r *http.Request) bool {
	return true
}

func CheckLogin(username, password string, r *http.Request) bool {
	return true
}
