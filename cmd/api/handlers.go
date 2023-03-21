package main

import (
	"encoding/json"
	"net/http"
)

type stripePayload struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
	Content string `json:"content"`
	ID      int    `json:"id"`
}

func (app *application) GetPaymentIntent(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK: true,
	}

	out, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		app.erroLog.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
