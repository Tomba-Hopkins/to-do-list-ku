package entities

import "time"

type User struct {
	Id   uint
	Username string
	Password string
}

type Todo struct {
	Id       uint
	Todoco   string
	Todoce   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Resol struct {
	IdR    uint
	Resolco string
	Resolce string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Note struct {
	Id   uint
	Note string
	Pencatat User
	CreatedAt time.Time
	UpdatedAt time.Time
}


// COBA DAN BISA : 
	// if r.Method == "POST"{
	// 	r.ParseForm()
	// 	username := r.FormValue("username")
	// 	password := r.FormValue("password")

	// 	if username == "kenzler" && password == "kenzler" || username == "kucing" && password == "kucing" {
	// 		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	// 		return
	// 	}
	// 	fmt.Fprintf(w,` <!DOCTYPE html>
	// 		<html>
	// 		  <head>
	// 			<title>hhhhhhh</title>
	// 		  </head>
	// 		  <body>
	// 			  <div style="width:100vw; height:100vh; background-color:black; color:white; display:flex; align-items:center; justify-content:center;">
	// 				<h1>LOE SIAPA JIR</h1>
	// 			  </div>
	// 		  </body>
	// 		</html>`)
	// 	return
	// }