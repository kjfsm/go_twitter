package main

import (
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-pg/pg"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	tDatabase := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "password",
		Database: "go_twitter",
	})
	defer tDatabase.Close()

	var tUsers []User
	tErr := tDatabase.Model(&tUsers).Select()
	if tErr != nil {
		panic(tErr)
	}

	tRouter := chi.NewRouter()
	tRouter.Get("/", func(tWriter http.ResponseWriter, r *http.Request) {
		tBuffer, _ := ioutil.ReadFile("users.html")
		tBody := string(tBuffer)
		tTemplate, _ := template.New("users").Parse(tBody)

		tData := struct {
			Title string
			Items []User
		}{
			Title: "Users",
			Items: tUsers,
		}

		tTemplate.Execute(tWriter, tData)
	})
	http.ListenAndServe(":8080", tRouter)
}
