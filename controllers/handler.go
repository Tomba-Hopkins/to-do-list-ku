package controllers

import (
	"html/template"
	"net/http"
	"strconv"
	"time"
	"to-do-list/entities"
	"to-do-list/models"

	"github.com/kataras/go-sessions/v3"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// cara 1------------
	models.Login(w, r)

	templ, err := template.ParseFiles("views/index.html")
	if err != nil {
		http.Error(w, "SUSAH BET JIR HHHHH HHH", http.StatusInternalServerError)
		return
	}
	templ.Execute(w, nil)
}


func RegisterHandler(w http.ResponseWriter, r *http.Request){
	models.Register(w, r)
	templ, err := template.ParseFiles("views/register/index.html")
	if err != nil {
		panic(err)
	}
	templ.Execute(w, nil)
}


func DashboardHandler(w http.ResponseWriter, r *http.Request) {

	todo := models.GetAll()

	session := sessions.Start(w, r)
	if len(session.GetString("username")) == 0 {
		http.Redirect(w, r, "/", 301)
	}


	if r.Method == "GET" {
		templ, err := template.ParseFiles("views/dashboard/index.html")
		if err != nil {
			panic(err)
	}
		data := map[string]any{
			"nama":session.GetString("username"),
			"todos":todo,
		}
		templ.Execute(w, data)
	}

	 // add todo
	 if r.Method == "POST" {
		var todo entities.Todo

		todo.Todoco = r.FormValue("todo1")
		todo.Todoce = r.FormValue("todo2")
		todo.CreatedAt = time.Now()
		todo.UpdatedAt = time.Now()

		if ok := models.CreateTodo(todo); !ok {
			templ, _ := template.ParseFiles("views/dashboard/index.html")
			templ.Execute(w, nil)
		}

		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	 }

}

func LogoutHandler(w http.ResponseWriter, r *http.Request){
	models.Logout(w, r)
}

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}
	if err := models.Delete(id); err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}