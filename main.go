package main

import (
	"log"
	"net/http"
	"to-do-list/config"
	"to-do-list/controllers"
	"to-do-list/models"
)

func main() {

	config.ConnectDb() // lupa ini jir

	mux := http.NewServeMux()

	defer config.DB.Close()


	// mux.HandleFunc("/styles/login", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "assets/style.css")
	// })
	// mux.HandleFunc("/styles/dashboard", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "assets/dashboard/style.css")
	// })
	// mux.HandleFunc("/script/dashboard", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "assets/dashboard/script.js")
	// })
	// mux.HandleFunc("/styles/register", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "assets/register/style.css")
	// })

	fileserver := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	// Path
	mux.HandleFunc("/", controllers.LoginHandler)
	mux.HandleFunc("/dashboard", controllers.DashboardHandler)
	mux.HandleFunc("/register", models.Register)
	mux.HandleFunc("/logout", controllers.LogoutHandler)
	mux.HandleFunc("/dashboard/delete", controllers.DeleteTodoHandler)

	log.Println("Starting Web at port 8080")
	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		panic(err)
	}

}