package controllers

import (
	"html/template"
	"net/http"
	"simple-login/models"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	errMsg := r.URL.Query().Get("error")
	tmpl := template.Must(template.ParseFiles("views/login.html"))
	tmpl.Execute(w, struct{ Error string }{errMsg})
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	user := r.FormValue("username")
	pass := r.FormValue("password")

	if !models.ValidateUser(user, pass) {
		http.Redirect(w, r, "/?error=Invalid+user+or+password", http.StatusSeeOther)
		return
	}

	sessionID := models.CreateSession(user)
	http.SetCookie(w, &http.Cookie{
		Name:  "session_id",
		Value: sessionID,
		Path:  "/",
	})
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err == nil {
		models.DeleteSession(cookie.Value)
	}
	http.SetCookie(w, &http.Cookie{
		Name:   "session_id",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
