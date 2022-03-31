package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)
}

func create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "create", nil)
}

func save_article(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	anons := r.FormValue("anons")
	full_text := r.FormValue("full_text")

	db, err := sql.Open("postgres", "user=postgres password=1603 dbname=articles sslmode=disable")
	CheckError(err)
	defer db.Close()
	fmt.Println(title, anons, full_text)
	insert, err := db.Query(`insert into "products"("title", "anons", "full_text") values($1, $2, $3)`, title, anons, full_text)
	CheckError(err)
	defer insert.Close()
	// insertDynStmt := `insert into "products"("title", "anons", "full_text") values($1, $2, $3)`
	// _, e = db.Exec(insertDynStmt, 2, "13 mini", "iPhone", 90000)
	// CheckError(e)

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func handleFunc() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/save_article", save_article)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleFunc()
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
