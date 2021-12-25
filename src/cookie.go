package main

import "net/http"

func cookieExistent(w http.ResponseWriter, r *http.Request, n string) bool {
	c, err := r.Cookie(n)
	if err != nil {
		return false
	}
	if c != nil {
		return true
	} else {
		return false
	}
}
