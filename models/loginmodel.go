package models

import (
	"net/http"
	"to-do-list/config"
	"to-do-list/entities"

	"github.com/kataras/go-sessions/v3"
	"golang.org/x/crypto/bcrypt"
)

var err error

func QueryUser(username string) entities.User {
	var users = entities.User{}
	 err = config.DB.QueryRow(
		`SELECT id, username, password FROM users WHERE username = ?`, username).Scan(&users.Id, &users.Username, &users.Password)

	return users
}


func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "views/register/index.html")
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	users := QueryUser(username)

	if (entities.User{}) == users {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		if len(hashedPassword) != 0 && err == nil {
			stmt, err := config.DB.Prepare("INSERT INTO users SET username = ?, password = ?")
			if err == nil {
				_, err := stmt.Exec(&username, &hashedPassword)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
		}
	} else {
		http.Redirect(w,r, "/register", 302)
	}
}


func Login(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)

	if len(session.GetString("username")) != 0 && err == nil {
		http.Redirect(w, r, "/dashboard", 302)
	}

	if r.Method != "POST" {
		http.ServeFile(w, r, "views/index.html")
		return
	}


	username := r.FormValue("username")
	password := r.FormValue("password")

	users := QueryUser(username)

	var passwordTest = bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(password))

	if passwordTest == nil {
		// success
		session := sessions.Start(w, r)
		session.Set("username", users.Username)
		http.Redirect(w, r, "/dashboard", 302)
	} else {
		// failed
		http.Redirect(w, r, "/", 302)
	}

}


func Logout(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	session.Clear()
	sessions.Destroy(w, r)
	http.Redirect(w, r, "/", 302)
}