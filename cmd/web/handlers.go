package main

import (
	"errors"
	"fmt"
	"github.com/euler-b/snippetbox/internal/models"

	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	snippet, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, snippet := range snippet {
		fmt.Fprintf(w, "%+v\n", snippet)
	}

	// estas lineas se comentan en este commit para poder hacer las consulta del listado de haikus...
	//
	//file := []string{
	//	"./ui/html/pages/base.tmpl.html",
	//	"./ui/html/partials/nav.tmpl.html",
	//	"./ui/html/pages/home.tmpl.html",
	//}
	//
	//th, err := template.ParseFiles(file...)
	//if err != nil {
	//	log.Print(err.Error())
	//	app.serverError(w, err)
	//	return
	//}
	//err = th.ExecuteTemplate(w, "base", nil)
	//if err != nil {
	//	log.Print(err.Error())
	//	app.serverError(w, err)
	//}

}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	fmt.Fprintf(w, "%+v", snippet)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)

}
