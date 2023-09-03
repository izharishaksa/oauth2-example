package main

import (
	"fmt"
	"net/http"
)

func authenticationPage(w http.ResponseWriter, r *http.Request) {
	originalURL := r.FormValue("original_url")
	if originalURL == "" {
		originalURL = "/"
	} else {
		originalURL = originalURL[1:]
	}

	token, _ := r.Cookie("auth_token")
	if token != nil && token.Value != "" {
		http.Redirect(w, r, originalURL, http.StatusFound)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<form action='/authenticate' method='POST'>")
	fmt.Fprintf(w, "<input type='text' name='username' placeholder='username' /><br/>")
	fmt.Fprintf(w, "<input type='password' name='password' placeholder='password' /><br/>")
	fmt.Fprintf(w, "<input type='hidden' name='original_url' value='%s' />", originalURL)
	fmt.Fprintf(w, "<input type='submit' value='Login' />")
	fmt.Fprintf(w, "</form>")
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if originalPassword, ok := users[username]; ok {
		if originalPassword == password {
			authToken, err := generateAuthToken(username)
			if err != nil {
				response(w, nil, err, http.StatusUnauthorized)
				return
			}

			http.SetCookie(w, &http.Cookie{
				Name:     "auth_token",
				Value:    authToken,
				HttpOnly: true,
			})

			originalURL := r.FormValue("original_url")
			http.Redirect(w, r, originalURL, http.StatusFound)
			return
		}
	}

	response(w, nil, fmt.Errorf("invalid username or password"), http.StatusUnauthorized)
}
