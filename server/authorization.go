package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func authorizationPage(w http.ResponseWriter, r *http.Request) {
	token, _ := r.Cookie("auth_token")
	if token == nil || token.Value == "" {
		http.Redirect(w, r, "/login?original_url="+url.QueryEscape(r.URL.String()), http.StatusFound)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<form action='/authorize' method='POST'>")
	fmt.Fprintf(w, "<input type='hidden' name='client_id' value='%s' />", r.FormValue("client_id"))
	fmt.Fprintf(w, "<input type='hidden' name='callback' value='%s' />", r.FormValue("callback"))
	fmt.Fprintf(w, "<input type='submit' value='Authorize' />")
	fmt.Fprintf(w, "</form>")
}

func authorize(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("auth_token")
	if err != nil {
		response(w, nil, fmt.Errorf(err.Error()), http.StatusUnauthorized)
		return
	}

	_, err = validateToken(token.Value)
	if err != nil {
		response(w, nil, fmt.Errorf(err.Error()), http.StatusUnauthorized)
		return
	}

	clientId := r.FormValue("client_id")
	client, ok := clients[clientId]
	if !ok {
		response(w, nil, fmt.Errorf("invalid client_id"), http.StatusUnauthorized)
		return
	}

	callback := r.FormValue("callback")
	if callback != client.CallbackUrl {
		response(w, nil, fmt.Errorf("invalid callback URL"), http.StatusUnauthorized)
		return
	}

	authCode, err := generateAuthCode(clientId)
	if err != nil {
		response(w, nil, err, http.StatusInternalServerError)
		return
	}

	authCodeToClientID[authCode] = clientId

	http.Redirect(w, r, fmt.Sprintf("%s?code=%s", callback, authCode), http.StatusFound)
}
