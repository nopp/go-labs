package controllers

import (
    "html/template"
    "net/http"
    "simple-login/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("session_id")
    if err != nil {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    username := models.GetUsernameFromSession(cookie.Value)
    if username == "" {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    tmpl := template.Must(template.ParseFiles("views/home.html"))
    tmpl.Execute(w, struct{ User string }{username})
}