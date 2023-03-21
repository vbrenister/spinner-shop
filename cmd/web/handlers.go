package main

import "net/http"

func (app *application) VeritualTerminal(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "terminal", nil); err != nil {
		app.erroLog.Println(err)
	}
}
