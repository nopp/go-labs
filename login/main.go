package main

import (
    "net/http"
    "simple-login/db"
    "simple-login/controllers"
)

func main() {
    if err := db.InitDB(); err != nil {
        panic(err)
    }
    defer db.DB.Close()
    db.CreateDefaultUser()

    http.HandleFunc("/", controllers.LoginPage)
    http.HandleFunc("/login", controllers.Login)
    http.HandleFunc("/home", controllers.Home)
    http.HandleFunc("/logout", controllers.Logout)

    http.ListenAndServe(":8080", nil)
}