package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func (app *appplication) home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	file := []string{
		"./ui/html/pages/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	th, err := template.ParseFiles(file...)
	if err != nil {
		log.Print(err.Error())
		app.serverError(w, err)
		return
	}
	err = th.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		app.serverError(w, err)
	}

}

func (app *appplication) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	fmt.Fprintf(w, "Aca tienes el snippet con el ID: %d", id)
}

func (app *appplication) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Esta feature permitira la creacion de un snippet"))
}
